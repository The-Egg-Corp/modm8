package paths

import (
	"os"
	"path/filepath"
)

func ConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, "modm8")
}

// Uses the users config dir and returns a path to the mod cache.
func ModCacheDir() string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, "modm8", "ModCache")
}

func NexusKeyPath() string {
	return filepath.Join(ConfigDir(), "nex.key")
}

func SettingsPath() string {
	return filepath.Join(ConfigDir(), "settings.toml")
}

func PersistencePath() string {
	return filepath.Join(ConfigDir(), "persistence.toml")
}
