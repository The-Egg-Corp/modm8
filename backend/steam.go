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

type SteamRunner struct {
	InstallPath *string
}

func NewSteamRunner() *SteamRunner {
	path, _ := GetSteamDirectory()

	return &SteamRunner{
		InstallPath: path,
	}
}

var homeDir string
var homeDirErr error

func (runner *SteamRunner) LaunchSteamGame(id uint32, args []string) (*exec.Cmd, error) {
	path, err := GetSteamDirectory()
	if err != nil {
		return nil, err
	}

	var platformExtension string
	if runtime.GOOS == "windows" {
		platformExtension = "Steam.exe"
	} else {
		platformExtension = "steam.sh"
	}

	// Construct a platform independent command that will launch the game with specified arguments.
	cmd := exec.Command(filepath.Join(*path, platformExtension), "-applaunch", strconv.Itoa(int(id)))
	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	return cmd, cmd.Run()
}

// TODO: Try use a global settings instance instead of creating a new one.
func GetSteamDirectory() (*string, error) {
	// Try and return path if it already exists in settings.toml
	settings := core.NewSettings()
	if err := settings.Load(); err != nil {
		return nil, fmt.Errorf("failed to load settings: %v", err)
	}

	if settings.General.SteamInstallPath != nil {
		return settings.General.SteamInstallPath, nil
	}

	dir, err := TryFindSteam()
	if err != nil {
		return nil, err
	}

	// Save to viper config & settings.toml file.
	// Prevents having to try find Steam again in future.
	settings.SetSteamInstallPath(*dir)
	settings.Save()

	return dir, nil
}

func TryFindSteam() (*string, error) {
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
				k.Close()
				continue
			}

			defer k.Close()
			installPath, _, err := k.GetStringValue("InstallPath")

			if err == nil {
				return &installPath, nil
			}
		}
	case "linux":
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
	}

	return nil, nil
}

func buildPath(args ...string) string {
	return filepath.Join(append([]string{homeDir}, args...)...)
}
