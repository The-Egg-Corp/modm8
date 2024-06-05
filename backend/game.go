package backend

import (
	"os"
	"path/filepath"
)

type StorePlatform string

const (
	STEAM            StorePlatform = "Steam"
	STEAM_DIRECT     StorePlatform = "Steam "
	EPIC_GAMES_STORE StorePlatform = "Epic Games Store"
	OCULUS_STORE     StorePlatform = "Oculus Store"
	ORIGIN           StorePlatform = "Origin / EA Desktop"
	XBOX_GAME_PASS   StorePlatform = "Xbox Game Pass"
	OTHER            StorePlatform = "Other"
)

type Game struct {
	DisplayName        string   `json:"display_name" mapstructure:"display_name"`
	SteamFolderName    string   `json:"steam_folder_name" mapstructure:"steam_folder_name"`
	SettingsIdentifier string   `json:"settings_identifier" mapstructure:"settings_identifier"`
	ExeName            []string `json:"exe_name" mapstructure:"exe_name"`
	DataFolderName     string   `json:"data_folder_name" mapstructure:"data_folder_name"`
	ThunderstoreURL    *string  `json:"thunderstore_url" mapstructure:"thunderstore_url"`
	ExclusionsURL      string   `json:"exclusions_url" mapstructure:"exclusions_url"`
}

type GameManager struct {
}

func NewGameManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) BepinexInstalled(absPath string) bool {
	installed, _ := BepinexInstalled(absPath)
	return installed
}

func BepinexInstalled(absPath string) (bool, []string) {
	required := []string{
		"BepInEx",
		"BepInEx/core",
		"BepInEx/core/BepInEx.dll",
		"BepInEx/core/BepInEx.Preloader.dll",
	}

	missing := []string{}

	for _, file := range required {
		// Use abs path to get path to current required file.
		path := filepath.Join(absPath, file)

		// Check if file exists, add to missing slice if not.
		if _, err := os.Stat(path); os.IsNotExist(err) {
			missing = append(missing, file)
		}
	}

	// If no missing files, BepInEx is likely installed.
	return len(missing) == 0, missing
}
