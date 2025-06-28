package backend

import (
	"modm8/backend/common/profile"
	"modm8/backend/game"
	"path/filepath"
	"strings"
	"testing"

	"github.com/the-egg-corp/thundergo/util"
)

// Requires creating a profile called "test".
func TestBepinexInstalled(t *testing.T) {
	testProfPath := profile.PathToProfile("Lethal Company", "test")
	installed, missing := game.BepinexInstalled(testProfPath)
	if !installed {
		for i, item := range missing {
			missing[i] = "  - " + item
		}

		missingStr := strings.Join(missing, "\n")
		t.Errorf("\n\nbepinex not installed at \"%s\"\nMissing items:\n%s", testProfPath, missingStr)
	}
}

// Requires IntroTweaks (or another mod) cfg to exist in a profile called "test".
func TestParseBepinexConfig(t *testing.T) {
	cfgPath := filepath.Join(profile.GameProfilesPath("Lethal Company"), "test", "BepInEx", "config", "IntroTweaks.cfg")
	parsed, err := game.ParseBepinexConfig(cfgPath)
	if err != nil {
		t.Fatal(err)
		return
	}

	util.PrettyPrint(parsed)
}

const testBepinexPackURL = "https://thunderstore.io/package/download/BepInEx/BepInExPack/5.4.2100/"

func TestInstallBepinexPack(t *testing.T) {
	profDir := profile.PathToProfile("Lethal Company", "test")

	_, err := game.InstallBepinexPack(testBepinexPackURL, profDir)
	if err != nil {
		t.Fatalf("failed to install BepInEx-Pack:\n%s", err)
	}
}
