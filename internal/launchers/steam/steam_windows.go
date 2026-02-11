//go:build windows

package steam

import (
	"golang.org/x/sys/windows/registry"
)

const platformExecName = "Steam.exe"

var regPaths = [...]string{
	`Software\Valve\Steam`,
	`Software\WOW6432Node\Valve\Steam`,
}

// Checks Windows Registry for the Steam installation path in both 32 and 64 bit paths.
func TryFindSteam() (*string, error) {
	for _, path := range regPaths {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		defer key.Close()

		installPath, _, err := key.GetStringValue("InstallPath")
		return &installPath, err
	}

	return nil, nil
}
