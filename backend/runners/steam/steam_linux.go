//go:build linux

package steam

import (
	"os"
	"path/filepath"
)

const platformExtension = "steam.sh"

var homeDir string
var homeDirErr error

func buildPath(args ...string) string {
	return filepath.Join(append([]string{homeDir}, args...)...)
}

func TryFindSteam() (*string, error) {
	homeDir, homeDirErr = os.UserHomeDir()
	if homeDirErr != nil {
		return nil, homeDirErr
	}

	paths := []string{
		buildPath(".steam"),
		buildPath(".steam", "steam"),
		buildPath(".steam", "root"),
		buildPath(".local", "share", "Steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam", "steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam", "root"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam"),
	}

	// Check all known steam paths for a shell file.
	for _, path := range paths {
		_, err := os.Stat(filepath.Join(path, "steam.sh"))

		if os.IsExist(err) {
			return &path, nil
		}

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
