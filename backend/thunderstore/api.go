package thunderstore

import (
	"context"
	"errors"
	"fmt"
	"modm8/backend"
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/samber/lo"

	exp "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
	v1 "github.com/the-egg-corp/thundergo/v1"
)

var modExceptions = []string{
	"Foldex-r2mod_cli",
	"ebkr-r2modman",
	"ebkr-r2modman_dsp",
	"ebkr-BT2TS",
	"ethanbrews-RiskOfRainModManager",
	"ethanbrews-Forecast_Mod_Manager",
	"scottbot95-RoR2ModManager",
	"HoodedDeath-RiskOfDeathModManager",
	"MythicManiac-MythicModManager",
	"Kesomannen-GaleModManager",
	"Elaviers-GCManager",
	"MADH95Mods-JSONRenameUtility",
	"Higgs1-Lighthouse",
}

var CurModCacheDir string

func ModCacheDir(gameTitle string) string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, "modm8", "Games", gameTitle, "ModCache")
}

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
}

func NewAPI(ctx context.Context) *API {
	return &API{
		Ctx:   ctx,
		Cache: NewCache(),
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

func (a *API) getCachedPackageList(community string) (v1.PackageList, error) {
	pkgs, exists := a.Cache[community]
	if !exists {
		return nil, errors.New("specified community has not been cached")
	}

	return pkgs, nil
}

func (a *API) GetCachedPackages(community string) ([]v1.Package, error) {
	return a.getCachedPackageList(community)
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

// Get packages for a specific community given its identifier.
// If the cache is not hit, it will be populated automatically.
func (a *API) GetPackagesInCommunity(community string, skipCache bool) ([]v1.Package, error) {
	if !skipCache {
		pkgs, exists := a.Cache[community]
		if exists {
			return pkgs, nil
		}
	}

	pkgs, err := v1.PackagesFromCommunity(community)
	if err != nil {
		return nil, err
	}

	a.Cache[community] = pkgs
	return pkgs, nil
}

// Similar to GetPackagesInCommunity, but every package is stripped of some info,
// massively decreasing the time spent sending data to the frontend to preventing it from blocking.
// The following keys are stripped (though retained in the cache):
//
// # Versions, DonationLink, Pinned
func (a *API) GetStrippedPackages(community string, skipCache bool) ([]StrippedPackage, error) {
	pkgs, err := a.GetPackagesInCommunity(community, skipCache)
	if err != nil {
		return nil, err
	}

	strippedPkgs := make([]StrippedPackage, 0, len(pkgs))

	// Loops over all pkgs, stripping some unnecessary fields to avoid blocking frontend.
	for _, pkg := range pkgs {
		// Strip any apps/utils that aren't strictly mods.
		if backend.ContainsEqualFold(modExceptions, pkg.FullName) {
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

func (a *API) GetPackagesByUser(communities []string, owner string) string {
	pkgs, err := v1.PackagesFromCommunities(v1.NewCommunityList(communities...))
	if err != nil {
		return "An error occurred getting packages!"
	}

	pkgs = lo.Filter(pkgs, func(pkg v1.Package, index int) bool {
		return strings.EqualFold(pkg.Owner, owner)
	})

	pkgCount := pkgs.Size()
	if pkgCount == 1 {
		return pkgs[0].Name
	}

	names := make([]string, pkgCount)
	for i := range pkgCount {
		names = append(names, pkgs[i].Name)
	}

	return strings.Join(names, ", ")
}

// Using the given list of packages (usually from a community), this function attempts to download the given
// package aka 'pkg' and any packages it depends on which is gathered by its 'Dependencies' field - see [PackageVersion].
//
// This function is recursive and calls Install for each dependency, any errors are accumulated into a slice and
// the install count is incremented if no error occurred - both of which are available once this function is complete.
func InstallWithDependencies(pkg v1.PackageVersion, pkgs v1.PackageList, errs *[]error, installCount *int) {
	_, err := Install(pkg, CurModCacheDir)
	if err == nil {
		*installCount += 1
	}

	for _, dependency := range pkg.Dependencies {
		// Split the dependency string into 3 elements: Author, Package Name, Version
		split := strings.Split(dependency, "-")
		pkg := pkgs.Get(split[0], split[1])

		if pkg == nil {
			*errs = append(*errs, fmt.Errorf("dependency '%s' not found", dependency))
			continue
		}

		InstallWithDependencies(*pkg.GetVersion(split[2]), pkgs, errs, installCount)
	}
}

// Downloads the given package version as a zip and unpacks it to the specified directory (expected to be absolute).
func Install(pkg v1.PackageVersion, dir string) (*grab.Response, error) {
	if exists, _ := backend.ExistsInDir(dir, pkg.FullName); exists {
		return nil, fmt.Errorf("%s is already installed", pkg.FullName)
	}

	resp, err := downloader.DownloadZip(pkg.DownloadURL, dir, pkg.FullName)
	if err != nil {
		return resp, err
	}

	// Finished, return if any error occurred.
	if err = resp.Err(); err != nil {
		return resp, err
	}

	path := filepath.Join(dir, pkg.FullName)
	ext := downloader.CUSTOM_ZIP_EXT

	// TODO: If the program closes for any reason, we need to be able to cancel (and possibly resume)
	// 		 installing the current zip, then also ensure it is deleted. Maybe when user next opens app?

	// Extract zip contents at path and delete the zip when done.
	err = fileutil.Unzip(path+ext, path, true)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Installs the latest version of a package by its full name (Owner-PkgName) and all of its dependencies.
//
// The game identifier (aka community) must be correctly specified for the package to be found.
func (a *API) InstallByName(gameTitle, community, fullName string) (*v1.PackageVersion, error) {
	// Get cached package list, if empty, try to fill it.
	pkgs, _ := a.getCachedPackageList(community)
	if pkgs == nil {
		commPkgs, err := v1.PackagesFromCommunity(community)
		if err != nil {
			return nil, fmt.Errorf("error getting packages: %s", err)
		}

		a.Cache[community] = commPkgs
	}

	pkg := pkgs.GetExact(fullName)
	if pkg == nil {
		return nil, fmt.Errorf("could not find package in community")
	}

	var errs []error
	var downloadCount int

	CurModCacheDir = ModCacheDir(gameTitle)

	latestVer := pkg.LatestVersion()
	InstallWithDependencies(latestVer, a.Cache[community], &errs, &downloadCount)

	return &latestVer, nil
}

// Downloads the specified package as a zip file and unpacks it under the specified directory (absolute path).
//
// The `fullName` parameter expects a string in the format: "Author-Package-Major.Minor.Patch"
// func (a *API) InstallPackage(fullName, dir string) (*grab.Response, error) {
// 	url := "https://thunderstore.io/package/download/" + strings.ReplaceAll(fullName, "-", "/")

// 	resp, err := downloader.DownloadZip(url, dir, fullName)
// 	if err != nil {
// 		return resp, err
// 	}

// 	// 	ticker := time.NewTicker(5 * time.Millisecond)
// 	// 	defer ticker.Stop()

// 	// Loop:
// 	// 	for {
// 	// 		select {
// 	// 		case <-resp.Done:
// 	// 			//timeTaken = time.Since(startTime)
// 	// 			break Loop
// 	// 		case <-ticker.C:
// 	// 			wRuntime.EventsEmit(a.Ctx, resp.Filename, 100*resp.Progress())
// 	// 		}
// 	// 	}

// 	// Finished, check for errors.
// 	if err = resp.Err(); err != nil {
// 		return resp, err
// 	}

// 	path := filepath.Join(dir, fullName)

// 	err = fileutil.Unzip(path+".zip", path, true)
// 	if err != nil {
// 		return resp, err
// 	}

// 	return resp, nil
// }
