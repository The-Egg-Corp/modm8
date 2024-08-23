package game

import (
	"path/filepath"

	backend "modm8/backend"
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
	NexusURL           *string  `json:"nexus_url" mapstructure:"nexus_url"`
	ExclusionsURL      string   `json:"exclusions_url" mapstructure:"exclusions_url"`
}

type GameManager struct{}

func NewManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) GameInstalled(dirPath string, exeKeywords []string) bool {
	var installed = false

	for _, keyword := range exeKeywords {
		path := filepath.Join(dirPath, keyword)
		if exists, _ := backend.ExistsAtPath(path); exists {
			installed = true
			break
		}
	}

	return installed
}

func (gm *GameManager) BepinexInstalled(absPath string) bool {
	installed, _ := BepinexInstalled(absPath)
	return installed
}

func (gm *GameManager) BepinexConfigFiles(dirs []string) ([]string, error) {
	return backend.WalkDirExt(filepath.Join(dirs...), []string{"cfg"})
}

func (gm *GameManager) ParseBepinexConfig(path string) (*BepinexConfig, error) {
	return ParseBepinexConfig(path)
}
