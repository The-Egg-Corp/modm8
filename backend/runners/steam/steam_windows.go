//go:build windows

package steam

import (
	"golang.org/x/sys/windows/registry"
)

const platformExecName = "Steam.exe"

var regPaths = []string{
	`Software\Valve\Steam`,
	`Software\WOW6432Node\Valve\Steam`,
}

func TryFindSteam() (*string, error) {
	// Check Windows Registry for Steam installation path in both 32 and 64 bit paths.
	for _, path := range regPaths {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
		if err != nil {
			continue
		}

		installPath, _, err := key.GetStringValue("InstallPath")
		key.Close()

		if err == nil {
			return &installPath, nil
		}
	}

	return nil, nil
}
