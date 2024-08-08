package steam

import (
	"fmt"
	"modm8/backend/app"
	"os/exec"
	"path/filepath"
	"strconv"
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
func (runner SteamRunner) LaunchSteamGame(id uint32, args []string) error {
	_, err := LaunchGame(id, args)
	return err
}

func LaunchGame(id uint32, args []string) (*exec.Cmd, error) {
	path, err := GetInstallDirectory()
	if err != nil {
		return nil, err
	}

	// Construct a platform independent command that will launch the game with specified arguments.
	cmd := exec.Command(filepath.Join(*path, platformExtension), "-applaunch", strconv.Itoa(int(id)))
	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	return cmd, cmd.Run()
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
