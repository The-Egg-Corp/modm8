package backend

import (
	"fmt"
	"modm8/backend/core"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

var homeDir string
var homeDirErr error

func buildPath(args ...string) string {
	return filepath.Join(append([]string{homeDir}, args...)...)
}

func linuxPathList() ([]string, error) {
	homeDir, homeDirErr = os.UserHomeDir()
	if homeDirErr != nil {
		return nil, homeDirErr
	}

	return []string{
		buildPath(".steam"),
		buildPath(".steam", "steam"),
		buildPath(".steam", "root"),
		buildPath(".local", "share", "Steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam", "steam"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".steam", "root"),
		buildPath(".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam"),
	}, nil
}

// TODO: Try use a global settings instance instead of creating a new one.
func FindSteamDirectory() (*string, error) {
	// Try and return path if it already exists in settings.toml
	settings := core.NewSettings()
	if err := settings.Load(); err != nil {
		return nil, fmt.Errorf("failed to load settings: %v", err)
	}

	if settings.General.SteamInstallPath != nil {
		return settings.General.SteamInstallPath, nil
	}

	// Otherwise try find where it might be from a list of paths.
	switch runtime.GOOS {
	case "windows":
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

			defer k.Close()
			installPath, _, err := k.GetStringValue("InstallPath")

			if err == nil {
				return &installPath, nil
			}
		}
	case "linux":
		paths, err := linuxPathList()
		if err != nil {
			return nil, err
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
	}

	return nil, nil
}

func LaunchSteamGame(id uint32, args ...string) error {
	path, err := FindSteamDirectory()
	if err != nil {
		return err
	}

	var platformExtension string
	if runtime.GOOS == "windows" {
		platformExtension = "Steam.exe"
	} else {
		platformExtension = "steam.sh"
	}

	// Construct the command to launch the game via Steam using -applaunch
	cmd := exec.Command(filepath.Join(*path, platformExtension), "-applaunch", strconv.Itoa(int(id)))
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
