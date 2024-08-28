package steam

import (
	"fmt"
	"modm8/backend/app"
	"path/filepath"
	"strconv"

	gocmd "github.com/go-cmd/cmd"
)

type SteamRunner struct {
	InstallPath *string
}

func NewRunner() *SteamRunner {
	path, _ := GetInstallDirectory()

	return &SteamRunner{
		InstallPath: path,
	}
}

// Attach to struct so Wails is aware of it.
func (runner SteamRunner) LaunchGame(id uint32, args []string) error {
	_, err := LaunchGame(id, args)
	return err
}

func LaunchGame(id uint32, args []string) (*gocmd.Cmd, error) {
	installDir, err := GetInstallDirectory()
	if err != nil {
		return nil, err
	}

	// Platform independent cmd that requests Steam to launch a game by its ID with all given arguments.
	// When constructed, it will look something like: "C:/path/to/steam -applaunch 69 --example_arg true"
	cmd := gocmd.NewCmd(filepath.Join(*installDir, platformExtension), "-applaunch", strconv.Itoa(int(id)))
	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	statusChan := cmd.Start()
	status := <-statusChan

	// Clean that bitch up properly.
	// See -> https://github.com/go-cmd/cmd?tab=readme-ov-file#proper-process-termination
	err = cmd.Stop()

	if status.Error != nil {
		err = status.Error
	}

	return cmd, err
}

// TODO: Try use a global settings instance instead of creating a new one.
func GetInstallDirectory() (*string, error) {
	// Try and return path if it already exists in settings.toml
	settings := app.NewSettings()
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
