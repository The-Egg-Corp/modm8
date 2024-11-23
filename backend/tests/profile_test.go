package backend

import (
	"fmt"
	"modm8/backend/common/profile"
	"testing"

	"github.com/the-egg-corp/thundergo/util"
)

const testGameTitle = "Lethal Company"

func TestGetProfileNames(t *testing.T) {
	names, err := profile.GetProfileNames(testGameTitle)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(names)
}

func TestGetProfiles(t *testing.T) {
	profs, err := profile.GetProfiles(testGameTitle)
	if err != nil {
		t.Fatal(err)
	}

	util.PrettyPrint(profs)
}

func TestSaveManifest(t *testing.T) {
	manifest := profile.NewProfileManifest()

	manifest.AddThunderstoreMod("Owen3H-IntroTweaks-1.5.0")
	manifest.AddThunderstoreMod("Owen3H-CSync-3.0.0")

	err := profile.SaveManifest(testGameTitle, "default", manifest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetManifest(t *testing.T) {
	_, err := profile.GetManifest(testGameTitle, "default")
	if err != nil {
		t.Fatal(err)
	}
}
