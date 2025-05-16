package profile

import (
	"errors"
	"slices"
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
