package backend

import (
	"modm8/backend"
	"strings"
	"testing"
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
