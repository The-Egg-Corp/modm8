package profile

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

func (manifest *ProfileManifest) AddThunderstoreMod(verFullName string) {
	manifest.Mods.Thunderstore = append(manifest.Mods.Thunderstore, verFullName)
}

func (manifest *ProfileManifest) AddNexusMod(verFullName string) {
	manifest.Mods.Nexus = append(manifest.Mods.Nexus, verFullName)
}
