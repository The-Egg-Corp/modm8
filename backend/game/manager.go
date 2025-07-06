package game

import (
	"modm8/backend/common/fileutil"
	"modm8/backend/common/profile"
	"modm8/backend/installing"
	"modm8/backend/loaders"
	"os"
	"path/filepath"
)

type GameManager struct{}

func NewGameManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) GetModInstaller(loader loaders.ModLoaderType) (installing.IModInstaller, error) {
	return installing.GetModInstaller(loader)
}

func (gm *GameManager) GetModLoader(loader loaders.ModLoaderType) (loaders.IModLoader, error) {
	return loaders.GetModLoader(loader)
}

func (gm *GameManager) GetModLinkPath(loader loaders.ModLoaderType, profileDir string) (string, error) {
	return loaders.GetModLinkPath(loader, profileDir)
}

// Creates a Symlink (uses Junction on Windows) in the respective loader's mod path and links it to a mod which must exist in the mod cache.
//
// This means that mods *technically* don't exist outside of the main mod cache of the game, mods in a profile are merely mirrors.
// This allows us to not only customise where the mod cache should be stored, but also have a single
// source of truth for mods, eliminating duplicates and saving disk space.
//
// For example, we can mirror target "../modm8/Games/GameTitle/Profiles/test/BepInEx/plugins/Owen3H-IntroTweaks-1.5.0" to the
// source "../modm8/Games/GameTitle/ModCache/Owen3H-IntroTweaks-1.5.0" which would give us the desired behaviour.
func (gm *GameManager) LinkModToProfile(loader loaders.ModLoaderType, gameTitle, profileName, modFullName string) error {
	profileModsDir, err := loaders.GetModLinkPath(loader, profile.PathToProfile(gameTitle, profileName))
	if err != nil {
		return err
	}

	// TODO: Read manifest.json in the mod's folder and call this function recursively for
	// 		 each mod specified within the `dependencies` field.

	target := filepath.Join(profileModsDir, modFullName)         // Path to mod in given profile dir
	source := filepath.Join(ModCacheDir(gameTitle), modFullName) // Path to mod in mod cache.

	return fileutil.LinkDir(target, source)
}

func (gm *GameManager) UnlinkModFromProfile(loader loaders.ModLoaderType, gameTitle, profileName, modFullName string) error {
	profileModsDir, err := loaders.GetModLinkPath(loader, profile.PathToProfile(gameTitle, profileName))
	if err != nil {
		return err
	}

	target := filepath.Join(profileModsDir, modFullName) // Path to mod in given profile dir
	return os.Remove(target)
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

// Uses the users config dir and returns a path to the mod cache for the input game.
// For example, input "Lethal Company" would return "%AppData%/Roaming/modm8/Games/Lethal Company/ModCache" on Windows.
//
// This function solely relys on `os` and `filepath` and should be platform-independent.
func ModCacheDir(gameTitle string) string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, "modm8", "Games", gameTitle, "ModCache")
}
