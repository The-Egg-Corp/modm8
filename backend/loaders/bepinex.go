package loaders

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Looks silly with no values, but allows constant lookup and avoids nested loop.
// An empty struct occupies no bytes in memory anyway :D
var PRELOADER_NAMES = map[string]struct{}{
	"BepInEx.Unity.Mono.Preloader.dll": {},
	"BepInEx.Unity.IL2CPP.dll":         {},
	"BepInEx.Preloader.dll":            {},
	"BepInEx.IL2CPP.dll":               {},
}

type BepinexLoaderInstructions struct {
}

func (instr BepinexLoaderInstructions) Generate(profileDir string) (*LoaderInstructions, error) {
	preloaderPath, err := GetBepinexPreloaderPath(profileDir)
	if err != nil {
		return nil, err
	}

	var instructions LoaderInstructions

	doorstopVer, _ := GetUnityDoorstopVersion(profileDir)
	if doorstopVer == 4 {
		instructions = GenDoorstopV4(preloaderPath)
	} else {
		// TODO: Should we assume V3 or just return an error like "unsupported doorstop ver"?
		instructions = GenDoorstopV3(preloaderPath)
	}

	return &instructions, nil
}

func GetBepinexPreloaderPath(profileDir string) (string, error) {
	coreDir := filepath.Join(profileDir, "BepInEx", "core")

	entries, err := os.ReadDir(coreDir)
	if err != nil {
		return "", fmt.Errorf("failed to read BepInEx core directory:\n%v", err)
	}

	// Search everything in this dir for a DLL matching one in PRELOADER_NAMES.
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue // Not a file, insta skip.
		}
		if _, ok := PRELOADER_NAMES[entry.Name()]; ok {
			return filepath.Join(coreDir, entry.Name()), nil
		}
	}

	return "", errors.New("BepInEx preloader not found")
}

func GenDoorstopV3(preloaderPath string) LoaderInstructions {
	return LoaderInstructions{
		ModdedParams: []string{
			"--doorstop-enable", "true",
			"--doorstop-target", preloaderPath,
		},
		VanillaParams: []string{
			"--doorstop-enable", "false",
		},
	}
}

func GenDoorstopV4(preloaderPath string) LoaderInstructions {
	return LoaderInstructions{
		ModdedParams: []string{
			"--doorstop-enabled", "true",
			"--doorstop-target-assembly", preloaderPath,
		},
		VanillaParams: []string{
			"--doorstop-enabled", "false",
		},
	}
}
