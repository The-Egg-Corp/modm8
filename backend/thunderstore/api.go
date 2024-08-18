package thunderstore

import (
	"errors"
	"strings"
	"sync"

	"github.com/samber/lo"
	exp "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
	v1 "github.com/the-egg-corp/thundergo/v1"
)

type API struct {
	Cache map[string]v1.PackageList
	mutex sync.RWMutex
}

func NewAPI() *API {
	return &API{
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

func (a *API) GetPackageVersions(community string, owner string, name string) ([]v1.PackageVersion, error) {
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
	// a.mutex.RLock()
	// defer a.mutex.RUnlock()

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

	// pkgs, err := comm.AllPackages()
	// if err != nil {
	// 	return make(map[string]v1.Package, 0), err
	// }

	pkgs, err := comm.AllPackages()
	if err != nil {
		return nil, err
	}

	// a.mutex.Lock()
	// defer a.mutex.Unlock()

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
