//go:build windows

package steam

import "golang.org/x/sys/windows/registry"

const platformExtension = "Steam.exe"

func TryFindSteam() (*string, error) {
	// Check Windows Registry for Steam installation path in both 32 and 64 bit paths.
	paths := []string{
		`Software\Valve\Steam`,
		`Software\WOW6432Node\Valve\Steam`,
	}

	for _, regPath := range paths {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, regPath, registry.QUERY_VALUE)
		if err != nil {
			continue
		}

		installPath, _, err := k.GetStringValue("InstallPath")
		k.Close()

		if err == nil {
			return &installPath, nil
		}
	}

	return nil, nil
}
