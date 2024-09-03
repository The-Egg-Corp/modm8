package thunderstore

import (
	"context"
	"errors"
	"fmt"
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cavaliergopher/grab/v3"
	"github.com/samber/lo"

	exp "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
	v1 "github.com/the-egg-corp/thundergo/v1"
)

type StrippedPackage struct {
	Name           string        `json:"name"`
	FullName       string        `json:"full_name"`
	Owner          string        `json:"owner"`
	UUID           string        `json:"uuid4"`
	PackageURL     string        `json:"package_url"`
	DateCreated    util.DateTime `json:"date_created"`
	DateUpdated    util.DateTime `json:"date_updated"`
	Rating         uint16        `json:"rating_score"`
	Deprecated     bool          `json:"is_deprecated"`
	HasNsfwContent bool          `json:"has_nsfw_content"`
	Categories     []string      `json:"categories"`
}

type API struct {
	Ctx   context.Context
	Cache map[string]v1.PackageList
	mutex sync.RWMutex
}

func NewAPI(ctx context.Context) *API {
	return &API{
		Ctx:   ctx,
		Cache: NewCache(),
		mutex: sync.RWMutex{},
	}
}

func NewCache() map[string]v1.PackageList {
	return make(map[string]v1.PackageList, 0)
}

func (a *API) ClearCache() {
	a.Cache = NewCache()
}

// Removes all packages associated with the specified community.
func (a *API) RemoveFromCache(community string) {
	delete(a.Cache, community)
}

func (a *API) GetCachedPackages(community string) ([]v1.Package, error) {
	pkgs, exists := a.Cache[community]
	if !exists {
		return nil, errors.New("specified community has not been cached")
	}

	return pkgs, nil
}

func (a *API) GetCommunities() (exp.CommunityList, error) {
	return exp.GetCommunities()
}

func (a *API) GetLatestPackageVersion(community string, owner string, name string) (*v1.PackageVersion, error) {
	versions, err := a.GetPackageVersions(community, owner, name)
	if err != nil {
		return nil, err
	}

	return &versions[0], nil
}

func (a *API) GetPackageVersions(community, owner, name string) ([]v1.PackageVersion, error) {
	pkgs, exists := a.Cache[community]
	if !exists {
		return nil, errors.New("specified community has not been cached")
	}

	pkg := pkgs.Get(owner, name)
	if pkg == nil {
		return nil, errors.New("no pkg exists with specified name")
	}

	return pkg.Versions, nil
}

func (a *API) GetPackagesInCommunity(community string, skipCache bool) ([]v1.Package, error) {
	if !skipCache {
		pkgs, exists := a.Cache[community]
		if exists {
			println("Cache hit")
			return pkgs, nil
		}
	}

	comm := v1.Community{
		Identifier: community,
	}

	pkgs, err := comm.AllPackages()
	if err != nil {
		return nil, err
	}

	a.Cache[community] = pkgs
	return pkgs, nil
}

func (a *API) GetPackagesStripped(community string, skipCache bool) ([]StrippedPackage, error) {
	pkgs, err := a.GetPackagesInCommunity(community, skipCache)
	if err != nil {
		return nil, err
	}

	strippedPkgs := make([]StrippedPackage, 0, len(pkgs))

	// Loops over all pkgs, stripping some unecessary fields, massively improving
	// time to serialize/deserialize to avoid blocking the frontend.
	for _, pkg := range pkgs {
		if strings.EqualFold(pkg.Name, "r2modman") {
			continue
		}

		strippedPkgs = append(strippedPkgs, StrippedPackage{
			Name:           pkg.Name,
			FullName:       pkg.FullName,
			Owner:          pkg.Owner,
			UUID:           pkg.UUID,
			PackageURL:     pkg.PackageURL,
			DateCreated:    pkg.DateCreated,
			DateUpdated:    pkg.DateUpdated,
			Rating:         pkg.Rating,
			Deprecated:     pkg.Deprecated,
			HasNsfwContent: pkg.HasNsfwContent,
			Categories:     pkg.Categories,
		})
	}

	return strippedPkgs, nil
}

func (a *API) GetUserPackages(communities []string, owner string) string {
	pkgs, err := v1.PackagesFromCommunities(v1.NewCommunityList(communities...))
	if err != nil {
		return "An error occurred getting packages!"
	}

	pkgs = lo.Filter(pkgs, func(pkg v1.Package, index int) bool {
		return strings.EqualFold(pkg.Owner, owner)
	})

	if pkgs.Size() == 1 {
		return pkgs[0].Name
	}

	var names []string
	var i uint

	for i = 0; i < pkgs.Size(); i++ {
		names = append(names, pkgs[i].Name)
	}

	return strings.Join(names, ", ")
}

// Returns the progress as a percentage
func (a *API) GetProgress(response grab.Response) float64 {
	return 100 * response.Progress()
}

// Downloads the specified package as a zip file and unpacks it under the specified directory (absolute path).
//
// The `fullName` parameter expects a string in the format: "Author-Package-Major.Minor.Patch"
func (a *API) InstallPackage(fullName, dir string) (*grab.Response, error) {
	url := "https://thunderstore.io/package/download/" + strings.ReplaceAll(fullName, "-", "/")

	resp, err := downloader.DownloadZip(url, dir, fullName)
	if err != nil {
		return resp, err
	}

	// 	ticker := time.NewTicker(5 * time.Millisecond)
	// 	defer ticker.Stop()

	// Loop:
	// 	for {
	// 		select {
	// 		case <-resp.Done:
	// 			//timeTaken = time.Since(startTime)
	// 			break Loop
	// 		case <-ticker.C:
	// 			wRuntime.EventsEmit(a.Ctx, resp.Filename, 100*resp.Progress())
	// 		}
	// 	}

	// Finished, check for errors.
	if err = resp.Err(); err != nil {
		return resp, err
	}

	path := filepath.Join(dir, fullName)

	err = fileutil.Unzip(path+".zip", path, true)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var testDir = GetProfileDir("LethalCompany", "test")

func GetProfileDir(game, profile string) string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, fmt.Sprintf("modm8\\Games\\%s\\Profiles\\%s\\BepInEx\\plugins", game, profile))
}

func InstallWithDependencies(ver v1.PackageVersion, pkgs v1.PackageList, errs *[]error, downloadCount *int) {
	_, err := Install(ver, testDir)
	if err == nil {
		*downloadCount += 1
	}

	for _, dependency := range ver.Dependencies {
		// Split the dependency string into 3 elements: Author, Package Name, Version
		split := strings.Split(dependency, "-")
		pkg := pkgs.Get(split[0], split[1])

		if pkg == nil {
			*errs = append(*errs, fmt.Errorf("dependency '%s' not found", dependency))
			continue
		}

		InstallWithDependencies(*pkg.GetVersion(split[2]), pkgs, errs, downloadCount)
	}
}

// Downloads the given package version as a zip and unpacks it to the specified directory (expected to be absolute).
func Install(pkg v1.PackageVersion, dir string) (*grab.Response, error) {
	resp, err := downloader.DownloadZip(pkg.DownloadURL, dir, pkg.FullName)
	if err != nil {
		return resp, err
	}

	// Finished, check for errors.
	if err = resp.Err(); err != nil {
		return resp, err
	}

	path := filepath.Join(dir, pkg.FullName)

	err = fileutil.Unzip(path+".zip", path, true)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
