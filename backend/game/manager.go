package game

import (
	"modm8/backend/common/fileutil"
	"modm8/backend/loaders"
	"os"
	"path/filepath"
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

func NewGameManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) GetLoaderInstructions(loader loaders.ModLoader, profileDir string) (*loaders.LoaderInstructions, error) {
	return loaders.GetLoaderInstructions(loader, profileDir)
}

// Creates a Symlink (uses Junction on Windows) at the target and links it to a mod which must exist in the cache.
// The target should be the path of the mod in the profile directory.
//
// This means that mods *technically* don't exist outside of the main mod cache of the game, mods in a profile are merely mirrors. This allows us to not only customise
// where the mod cache should be stored, but also have a single source of truth for mods, eliminating dupes and saving disk space.
//
// For example, we can mirror target "../modm8/Games/GameTitle/Profiles/test/BepInEx/plugins/Owen3H-IntroTweaks-1.5.0" to the
// source "../modm8/Games/GameTitle/ModCache/Owen3H-IntroTweaks-1.5.0" which would give us the desired behaviour.
func (gm *GameManager) LinkModToProfile(gameTitle string, target string) error {
	return fileutil.LinkDir(target, ModCacheDir(gameTitle))
}

func (gm *GameManager) GameInstalled(dirPath string, exeKeywords []string) bool {
	installed := false

	for _, keyword := range exeKeywords {
		path := filepath.Join(dirPath, keyword)
		if exists, _ := fileutil.ExistsAtPath(path); exists {
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
	return fileutil.GetFilesWithExts(filepath.Join(dirs...), []string{"cfg"})
}

func (gm *GameManager) ParseBepinexConfig(path string) (*BepinexConfig, error) {
	return ParseBepinexConfig(path)
}

func ModCacheDir(gameTitle string) string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, "modm8", "Games", gameTitle, "ModCache")
}
