package steam

import (
	"errors"
	"fmt"
	"modm8/backend/app"
	"path/filepath"
	"strconv"
	"strings"

	gocmd "github.com/go-cmd/cmd"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SteamRunner struct {
	// InstallPath    *string
	// InstallPathErr error
}

func NewSteamRunner() *SteamRunner {
	return &SteamRunner{}
}

// Attach to struct so Wails is aware of it.
func (runner SteamRunner) LaunchGame(id uint32, args []string) error {
	installDir, err := GetInstallDirectory()
	if err != nil {
		return err
	}

	_, err = LaunchGame(installDir, platformExecName, id, args)
	if err != nil {
		// Try again and look for "Steam" rather than "steam".
		ext := cases.Title(language.English, cases.NoLower).String(platformExecName)
		_, err = LaunchGame(installDir, ext, id, args)
	}

	return err
}

func LaunchGame(installDir *string, ext string, id uint32, args []string) (*gocmd.Cmd, error) {
	if installDir == nil {
		return nil, errors.New("cannot launch game. steam path must be non-nil")
	}
	if strings.TrimSpace(*installDir) == "" {
		return nil, errors.New("cannot launch game. steam path must be non-empty")
	}

	exePath := filepath.Join(*installDir, platformExecName)

	// Platform independent cmd that requests Steam to launch a game by its ID with all given arguments.
	// When constructed, it will look something like: "C:/path/to/steam -applaunch 69 --example_arg true"
	cmd := gocmd.NewCmd(exePath, "-applaunch", strconv.Itoa(int(id)))
	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	statusChan := cmd.Start()
	status := <-statusChan

	// Clean that bitch up properly.
	// See -> https://github.com/go-cmd/cmd?tab=readme-ov-file#proper-process-termination
	err := cmd.Stop()
	if status.Error != nil {
		err = status.Error
	}

	return cmd, err
}

// TODO: Try use a global settings instance instead of creating a new one.

// Returns the path to the directory where Steam is installed.
func GetInstallDirectory() (*string, error) {
	// Try and return path if it already exists in settings.toml
	settings := app.NewSettings()
	if err := settings.Load(); err != nil {
		return nil, fmt.Errorf("failed to load settings: %v", err)
	}

	settingsSteamPath := settings.Misc.SteamInstallPath
	if settingsSteamPath != nil && strings.TrimSpace(*settingsSteamPath) != "" {
		return settingsSteamPath, nil
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
