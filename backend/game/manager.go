package game

import (
	"path/filepath"

	backend "modm8/backend"
)

// type StorePlatform string

// const (
// 	STEAM            StorePlatform = "Steam"
// 	STEAM_DIRECT     StorePlatform = "Steam "
// 	EPIC_GAMES_STORE StorePlatform = "Epic Games Store"
// 	OCULUS_STORE     StorePlatform = "Oculus Store"
// 	ORIGIN           StorePlatform = "Origin / EA Desktop"
// 	XBOX_GAME_PASS   StorePlatform = "Xbox Game Pass"
// 	OTHER            StorePlatform = "Other"
// )

type GameManager struct{}

func NewManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) GameInstalled(dirPath string, exeKeywords []string) bool {
	installed := false

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
	return backend.GetFilesWithExts(filepath.Join(dirs...), []string{"cfg"})
}

func (gm *GameManager) ParseBepinexConfig(path string) (*BepinexConfig, error) {
	return ParseBepinexConfig(path)
}
