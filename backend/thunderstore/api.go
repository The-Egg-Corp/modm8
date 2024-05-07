package thunderstore

import (
	"strings"

	"github.com/samber/lo"
	exp "github.com/the-egg-corp/thundergo/experimental"
	v1 "github.com/the-egg-corp/thundergo/v1"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) GetCommunities() (exp.CommunityList, error) {
	return exp.GetCommunities()
}

func (a *API) GetPackagesInCommunity(identifier string) (v1.PackageList, error) {
	comm := v1.Community{
		Identifier: identifier,
	}

	return comm.AllPackages()
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
