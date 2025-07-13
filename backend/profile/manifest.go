package profile

import (
	"fmt"
	"modm8/backend/platform"
	"slices"
	"strings"

	"github.com/samber/lo"
)

const manifestName = "profinfo.json"

type ProfileMods map[platform.ModPlatform][]string
type ProfileManifest struct {
	Mods ProfileMods
}

func NewProfileManifest() ProfileManifest {
	return ProfileManifest{
		Mods: ProfileMods{
			platform.NEXUS:        []string{},
			platform.THUNDERSTORE: []string{},
		},
	}
}

func (manifest *ProfileManifest) AddMod(platform platform.ModPlatform, verFullName string) error {
	if !slices.Contains(manifest.Mods[platform], verFullName) {
		manifest.Mods[platform] = append(manifest.Mods[platform], verFullName)
	}

	// This should never happen in a perfect world, but this exists for the eventuality where
	// the frontend shits itself or some how, some way, the user is able to press install more than once.
	return fmt.Errorf("could not add mod from platform %s to profile manifest. already exists", platform)
}

func (manifest *ProfileManifest) RemoveMod(platform platform.ModPlatform, verFullName string) uint {
	prevLen := len(manifest.Mods[platform])
	manifest.Mods[platform] = lo.Filter(manifest.Mods[platform], func(el string, idx int) bool {
		return !strings.EqualFold(el, verFullName)
	})

	newLen := len(manifest.Mods[platform])
	return clampL(0, prevLen-newLen)
}

func clampL(min int, x int) uint {
	if x < min {
		x = min
	}

	return uint(x)
}
