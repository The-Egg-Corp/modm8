package profile

import (
	"errors"
	"slices"
	"strings"

	"github.com/samber/lo"
)

const manifestName = "profinfo.json"

type ProfileManifest struct {
	Mods ProfileMods `json:"mods"`
}

type ProfileMods struct {
	Thunderstore []string `json:"thunderstore"`
	Nexus        []string `json:"nexus"`
}

func NewProfileManifest() ProfileManifest {
	return ProfileManifest{
		Mods: ProfileMods{
			Thunderstore: []string{},
			Nexus:        []string{},
		},
	}
}

func (manifest *ProfileManifest) AddThunderstoreMod(verFullName string) error {
	if !slices.Contains(manifest.Mods.Thunderstore, verFullName) {
		manifest.Mods.Thunderstore = append(manifest.Mods.Thunderstore, verFullName)
	}

	// This should never happen in a perfect-world, but this exists for the eventuality where
	// the frontend shits itself or some how, some way, the user is able to press install more than once.
	return errors.New("could not add ts mod to profile manifest. already exists")
}

func (manifest *ProfileManifest) AddNexusMod(verFullName string) error {
	if !slices.Contains(manifest.Mods.Nexus, verFullName) {
		manifest.Mods.Nexus = append(manifest.Mods.Nexus, verFullName)
	}

	return errors.New("could not add nexus mod to profile manifest. already exists")
}

func (manifest *ProfileManifest) RemoveThunderstoreMod(verFullName string) uint {
	prev := len(manifest.Mods.Thunderstore)
	manifest.Mods.Thunderstore = lo.Filter(manifest.Mods.Thunderstore, func(el string, idx int) bool {
		return !strings.EqualFold(el, verFullName)
	})

	return clampL(0, prev-len(manifest.Mods.Thunderstore))
}

func (manifest *ProfileManifest) RemoveNexusMod(verFullName string) uint {
	prev := len(manifest.Mods.Nexus)
	manifest.Mods.Nexus = lo.Filter(manifest.Mods.Nexus, func(el string, idx int) bool {
		return !strings.EqualFold(el, verFullName)
	})

	return clampL(0, prev-len(manifest.Mods.Nexus))
}

func clampL(min int, x int) uint {
	if x < min {
		x = min
	}

	return uint(x)
}
