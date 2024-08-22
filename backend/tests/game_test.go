package backend

import (
	"modm8/backend"
	"strings"
	"testing"

	"github.com/the-egg-corp/thundergo/util"
)

func TestBepinexInstalled(t *testing.T) {
	dir := "E:\\SteamLibrary\\steamapps\\common\\Lethal Company"
	installed, missing := backend.BepinexInstalled(dir)

	if !installed {
		for i, item := range missing {
			missing[i] = "  - " + item
		}

		missingStr := strings.Join(missing, "\n")
		t.Errorf("\n\nbepinex not installed at \"%s\"\nMissing items:\n%s", dir, missingStr)
	}
}

func TestParseConfig(t *testing.T) {
	cfgPath := "E:\\SteamLibrary\\steamapps\\common\\Lethal Company\\BepInEx\\config\\IntroTweaks.cfg"
	parsed, err := backend.ParseBepinexConfig(cfgPath)

	if err != nil {
		t.Fatal(err)
		return
	}

	util.PrettyPrint(parsed)
}
