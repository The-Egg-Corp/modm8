package steam

import (
	"errors"
	"fmt"
	"modm8/backend/app/appctx"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/andygrunwald/vdf"
	gocmd "github.com/go-cmd/cmd"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SteamLauncher struct {
	appSettings *appctx.AppSettings
	// InstallPath    *string
	// InstallPathErr error
}

func NewSteamLauncher(appSettings *appctx.AppSettings) *SteamLauncher {
	return &SteamLauncher{appSettings: appSettings}
}

// Attach to struct so Wails is aware of it.
func (sl SteamLauncher) LaunchGame(id uint32, args []string) error {
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

func (sl SteamLauncher) GetLibPaths() ([]string, error) {
	return GetSteamLibPaths()
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

// Returns the path to the directory where Steam is installed.
func GetInstallDirectory() (*string, error) {
	// TODO: Instead of NewSettings(), use AppSettings from SteamLauncher in a way that is testable.
	settings := appctx.NewSettings()
	if err := settings.Load(); err != nil {
		return nil, fmt.Errorf("failed to load settings: %v", err)
	}

	// Try and return path if it already exists in ../modm8/settings.toml
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

func GetSteamLibVDF() (*string, error) {
	steamPath, err := GetInstallDirectory()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(*steamPath, "steamapps", "libraryfolders.vdf")
	return &path, nil
}

func GetSteamLibPaths() ([]string, error) {
	vdfPath, err := GetSteamLibVDF()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(*vdfPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m, err := vdf.NewParser(file).Parse()
	if err != nil {
		return nil, err
	}

	libFolders, ok := m["libraryfolders"].(map[string]any)
	if !ok {
		return nil, errors.New("failed to parse steam libraryfolders.vdf file: invalid structure")
	}

	paths := []string{}
	for _, raw := range libFolders {
		entry, ok := raw.(map[string]any) // Can't rly marshal since its not JSON.
		if !ok {
			continue
		}

		// We managed to get the path and cast it to a string.
		if path, ok := entry["path"].(string); ok {
			paths = append(paths, path)
		}
	}

	return paths, nil
}

// Searches through every SteamLibrary and consolidates each game path into single array like so:
//
//	"H:\Program Files (x86)\Steam\steamapps\common\Lethal Company",
//	"F:\SteamLibrary\steamapps\common\Risk of Rain 2"
func GetSteamGamePaths() ([]string, error) {
	libPaths, err := GetSteamLibPaths()
	if err != nil {
		return nil, err
	}

	allPaths := []string{}
	for _, lib := range libPaths {
		libPaths, err := GetGamesInSteamLib(lib)
		if err != nil {
			continue
		}

		allPaths = append(allPaths, libPaths...)
	}

	return allPaths, nil
}

func GetGamesInSteamLib(libPath string) ([]string, error) {
	common := filepath.Join(libPath, "steamapps", "common")
	entries, err := os.ReadDir(common)
	if err != nil {
		return nil, err
	}

	games := []string{}
	for _, e := range entries {
		if e.IsDir() {
			games = append(games, filepath.Join(common, e.Name()))
		}
	}

	return games, nil
}
