package backend

import (
	"modm8/backend/common/profile"
	"testing"
)

func TestWriteProfile(t *testing.T) {
	manifest := profile.ProfileManifest{
		Mods: profile.ProfileMods{
			Thunderstore: []string{"example-ts-mod-4.2.0"},
			Nexus:        []string{"example-nexus-mod-6.9.0"},
		},
	}

	err := profile.SaveManifest("Lethal Company", "test", manifest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadProfile(t *testing.T) {
	_, err := profile.GetManifest("Lethal Company", "test")
	if err != nil {
		t.Fatal(err)
	}
}
