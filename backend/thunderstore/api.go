package thunderstore

import (
	"context"
	"errors"
	"fmt"
	"modm8/backend/app"
	"modm8/backend/installing"
	"modm8/backend/loaders"
	"modm8/backend/utils"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/samber/lo"

	exp "github.com/the-egg-corp/thundergo/experimental"
	TSGOUtil "github.com/the-egg-corp/thundergo/util"
	TSGOV1 "github.com/the-egg-corp/thundergo/v1"
)

var modExclusions = []string{
	"ebkr-r2modman",
	"ebkr-r2modman_dsp",
	"ebkr-BT2TS",
	"Harb-AttributeFinder",
	"Kesomannen-GaleModManager",
	"ethanbrews-RiskOfRainModManager",
	"scottbot95-RoR2ModManager",
	"HoodedDeath-RiskOfDeathModManager",
	"Elaviers-GCManager",
	"MythicManiac-MythicModManager",
	"Higgs1-Lighthouse",
	"ethanbrews-Forecast_Mod_Manager",
	"3c079bcb4f34402f-BepInExPack",
	"Dasvcx-Enforcer",
	"SgtPopNFresh-BepInExPack",
	"TheLamiaLover-MobileTurretFungus",
	"Squidy-RogueWisp",
	"TheLamiaLover-ItemStatsMod",
	"Foldex-r2mod_cli",
	"gnonme-CustomItemsSDK",
	"gnonme-ModThatIsNotMod_Unity_Tools",
	"L4rs-QuickFSR",
	"MPModTeam-Boneworks_MP",
	"MADH95Mods-JSONRenameUtility",
	"GamefaceGamers-Mod_Sync",
	"Pyoid-Hook_Line_and_Sinker",
	"GardenGals-Hatchery",
}

// The dir where the mod cache is located for the current game.
var ModCacheDir = app.ModCacheDir()

// Same as a thundergo `Package` but the 'Versions' field is replaced with only a single 'LatestVersion' field
// and the following fields are completely removed: [DonationLink, Pinned].
type StrippedPackage struct {
	Name           string                `json:"name"`
	FullName       string                `json:"full_name"`
	Owner          string                `json:"owner"`
	UUID           string                `json:"uuid4"`
	PackageURL     string                `json:"package_url"`
	DateCreated    TSGOUtil.DateTime     `json:"date_created"`
	DateUpdated    TSGOUtil.DateTime     `json:"date_updated"`
	Rating         uint16                `json:"rating_score"`
	Deprecated     bool                  `json:"is_deprecated"`
	HasNsfwContent bool                  `json:"has_nsfw_content"`
	Categories     []string              `json:"categories"`
	LatestVersion  TSGOV1.PackageVersion `json:"latest_version"`
}

type ThunderstoreAPI struct {
	Ctx   context.Context
	Cache map[string]TSGOV1.PackageList // Cached packages for every community. Not related in any way to the ModCache dir.
}

func NewThunderstoreAPI(ctx context.Context) *ThunderstoreAPI {
	return &ThunderstoreAPI{
		Ctx:   ctx,
		Cache: NewCache(),
	}
}

func NewCache() map[string]TSGOV1.PackageList {
	return make(map[string]TSGOV1.PackageList, 0)
}

func (api *ThunderstoreAPI) ClearCache() {
	api.Cache = NewCache()
}

// Removes all packages associated with the specified community.
func (api *ThunderstoreAPI) RemoveFromCache(community string) {
	delete(api.Cache, community)
}

func (api *ThunderstoreAPI) getCachedPackageList(community string) (TSGOV1.PackageList, error) {
	pkgs, exists := api.Cache[community]
	if !exists {
		return nil, errors.New("specified community has not been cached")
	}

	return pkgs, nil
}

func (api *ThunderstoreAPI) GetCachedPackages(community string) ([]TSGOV1.Package, error) {
	return api.getCachedPackageList(community)
}

func (api *ThunderstoreAPI) GetCommunities() (exp.CommunityList, error) {
	return exp.GetCommunities()
}

func (api *ThunderstoreAPI) GetLatestPackageVersion(community string, owner string, name string) (*TSGOV1.PackageVersion, error) {
	versions, err := api.GetPackageVersions(community, owner, name)
	if err != nil {
		return nil, err
	}

	return &versions[0], nil
}

func (api *ThunderstoreAPI) GetPackageVersions(community, owner, name string) ([]TSGOV1.PackageVersion, error) {
	pkgs, exists := api.Cache[community]
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
func (api *ThunderstoreAPI) GetPackagesInCommunity(community string, skipCache bool) ([]TSGOV1.Package, error) {
	if !skipCache {
		pkgs, exists := api.Cache[community]
		if exists {
			return pkgs, nil
		}
	}

	pkgs, err := TSGOV1.PackagesFromCommunity(community)
	if err != nil {
		return nil, err
	}

	api.Cache[community] = pkgs
	return pkgs, nil
}

func (api *ThunderstoreAPI) GetPackagesByUser(communities []string, owner string) string {
	return GetPackagesByUser(communities, owner)
}

func GetPackagesByUser(communities []string, owner string) string {
	pkgs, err := TSGOV1.PackagesFromCommunities(TSGOV1.NewCommunityList(communities...))
	if err != nil {
		return "An error occurred getting packages!"
	}

	pkgs = lo.Filter(pkgs, func(pkg TSGOV1.Package, index int) bool {
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

// Similar to GetPackagesInCommunity, but every package is stripped of some info,
// massively decreasing the time spent sending data to the frontend to preventing it from blocking.
// The following keys are stripped (though retained in the cache):
//
// # Versions, DonationLink, Pinned
func (api *ThunderstoreAPI) GetStrippedPackages(community string, skipCache bool) ([]StrippedPackage, error) {
	pkgs, err := api.GetPackagesInCommunity(community, skipCache)
	if err != nil {
		return nil, err
	}

	strippedPkgs := make([]StrippedPackage, 0, len(pkgs))

	// Loops over all pkgs, stripping some unnecessary fields to avoid blocking frontend.
	for _, pkg := range pkgs {
		// Strip any apps/utils that aren't strictly mods.
		if utils.ArrEqualFold(modExclusions, pkg.FullName) {
			continue
		}

		strippedPkgs = append(strippedPkgs, StrippedPackage{
			Name:           pkg.Name,     // Ex: "CSync"
			FullName:       pkg.FullName, // Ex: "Owen3H-CSync"
			Owner:          pkg.Owner,
			UUID:           pkg.UUID,
			PackageURL:     pkg.PackageURL,
			DateCreated:    pkg.DateCreated,
			DateUpdated:    pkg.DateUpdated,
			Rating:         pkg.Rating,
			Deprecated:     pkg.Deprecated,
			HasNsfwContent: pkg.HasNsfwContent,
			Categories:     pkg.Categories,
			LatestVersion:  pkg.LatestVersion(),
		})
	}

	return strippedPkgs, nil
}

// Installs the latest version of a package by its full name (Owner-PkgName) and all of its dependencies.
//
// The game identifier (aka community) must be correctly specified for the package to be found.
func (api *ThunderstoreAPI) InstallByName(gameTitle, commIdent, fullName string) (*TSGOV1.PackageVersion, error) {
	ecosys, err := GetSchema().GetEcosystem()
	if err != nil {
		return nil, fmt.Errorf("could not get Thunderstore ecosystem")
	}

	r2mapping := ecosys.Games[commIdent].R2Modman[0]
	loader := loaders.GetModLoaderType(r2mapping.PackageLoader)

	// Get cached package list, if empty, try to fill it.
	pkgs, _ := api.getCachedPackageList(commIdent)
	if pkgs == nil {
		commPkgs, err := TSGOV1.PackagesFromCommunity(commIdent)
		if err != nil {
			return nil, fmt.Errorf("error getting packages: %s", err)
		}

		api.Cache[commIdent] = commPkgs
	}

	pkg := pkgs.GetExact(fullName)
	if pkg == nil {
		return nil, fmt.Errorf("could not find package in community")
	}

	var errs []error
	var downloadCount int

	latestVer := pkg.LatestVersion()
	meta := installing.PackageInstallMeta{
		Loader:       loader,
		FullName:     latestVer.FullName,
		DownloadURL:  latestVer.DownloadURL,
		Dependencies: latestVer.Dependencies,
	}

	InstallWithDependencies(meta, api.Cache[commIdent], &errs, &downloadCount)

	if len(errs) > 0 {
		var errBuilder strings.Builder
		for _, err := range errs {
			errBuilder.WriteString(err.Error())
			errBuilder.WriteString("\n")
		}

		return &latestVer, fmt.Errorf("errors occurred installing %s:\n%s", latestVer.FullName, errBuilder.String())
	}

	return &latestVer, nil
}

// Using the given list of packages (usually from a community), this function attempts to download the given
// package aka 'pkg' and any packages it depends on which is gathered by its 'Dependencies' field - see [v1.PackageVersion].
//
// This function is recursive and calls [Install] for each dependency, any errors are accumulated into a slice and
// the install count is incremented if no error occurred - both of which are available once this func has fully finished.
func InstallWithDependencies(pkgInsMeta installing.PackageInstallMeta, pkgs TSGOV1.PackageList, errs *[]error, installCount *int) {
	_, err := Install(pkgInsMeta, ModCacheDir)
	if err == nil {
		*installCount += 1
	}

	// Install dependencies of dependencies and so forth, until no more left.
	for _, dependency := range pkgInsMeta.Dependencies {
		// Split the dependency string into 3 elements: Author, Package Name, Version
		split := strings.Split(dependency, "-")

		pkg := pkgs.Get(split[0], split[1])
		if pkg == nil {
			*errs = append(*errs, fmt.Errorf("dependency '%s' not found", dependency))
			continue
		}

		ver := pkg.GetVersion(split[2])
		meta := installing.PackageInstallMeta{
			Loader:       pkgInsMeta.Loader,
			FullName:     ver.FullName,
			DownloadURL:  ver.DownloadURL,
			Dependencies: ver.Dependencies,
		}

		InstallWithDependencies(meta, pkgs, errs, installCount)
	}
}

// Downloads the given package version as a zip and unpacks it to the specified dir path (expected to be absolute).
func Install(pkgInsMeta installing.PackageInstallMeta, dirPath string) (*grab.Response, error) {
	ins, err := installing.GetModInstaller(pkgInsMeta.Loader)
	if err != nil {
		return nil, err
	}

	return ins.Install(pkgInsMeta.DownloadURL, pkgInsMeta.FullName, dirPath)
}

// Downloads the specified package as a zip file and unpacks it under the specified directory (absolute path).
// Uses a ticker to emit an event indicating the download progress every X milliseconds.
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
// 	// 			wuntime.EventsEmit(a.Ctx, resp.Filename, 100*resp.Progress())
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
