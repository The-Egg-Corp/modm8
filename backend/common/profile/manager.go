package profile

import (
	"encoding/json"
	"modm8/backend"
	"os"
	"path/filepath"
	"strings"
)

// type ModStore map[string]IMod
// type ProfileStore = map[string]Profile

type ProfileManager struct{}

func NewManager() *ProfileManager {
	return &ProfileManager{}
}

type ProfileMods struct {
	Thunderstore []string `json:"thunderstore"`
	Nexus        []string `json:"nexus"`
}

type ProfileManifest struct {
	Mods ProfileMods `json:"mods"`
}

func (pm *ProfileManager) GetProfiles(gameTitle string) (map[string]ProfileManifest, error) {
	return GetProfiles(gameTitle)
}

func (pm *ProfileManager) GetProfile(gameTitle, profileName string) (*ProfileManifest, error) {
	return GetManifest(gameTitle, profileName)
}

func (pm *ProfileManager) SaveProfile(gameTitle, profileName string, prof ProfileManifest) error {
	return SaveManifest(gameTitle, profileName, prof)
}

func PathToProfilesDir(gameTitle string) string {
	cacheDir, _ := os.UserConfigDir()
	path := filepath.Join(cacheDir, "modm8", "Games", gameTitle, "Profiles")

	return path
}

func PathToProfile(gameTitle, profileName string) string {
	return filepath.Join(PathToProfilesDir(gameTitle), profileName+".prof")
}

func GetProfileNames(gameTitle string) ([]string, error) {
	profDir := PathToProfilesDir(gameTitle)

	// The user probably hasn't created a profiles yet.
	exists, _ := backend.ExistsAtPath(profDir)
	if !exists {
		return []string{}, nil
	}

	paths, err := backend.WalkDirExt(profDir, []string{"prof"})
	if err != nil {
		return []string{}, err
	}

	var names []string
	for _, path := range paths {
		name := strings.Replace(filepath.Base(path), ".prof", "", -1)
		names = append(names, name)
	}

	return names, nil
}

func SaveManifest(gameTitle, profileName string, prof ProfileManifest) error {
	data, err := json.Marshal(prof)
	if err != nil {
		return err
	}

	return backend.SaveFile(PathToProfile(gameTitle, profileName), data)
}

func GetManifest(gameTitle, profileName string) (*ProfileManifest, error) {
	contents, err := backend.ReadFile(PathToProfile(gameTitle, profileName))
	if err != nil {
		return nil, err
	}

	var manifest ProfileManifest
	if err := json.Unmarshal(contents, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

func GetProfiles(gameTitle string) (map[string]ProfileManifest, error) {
	profNames, err := GetProfileNames(gameTitle)
	if err != nil {
		return nil, err
	}

	profiles := make(map[string]ProfileManifest)
	for _, name := range profNames {
		manifest, _ := GetManifest(gameTitle, name)
		if manifest != nil {
			profiles[name] = *manifest
		}
	}

	return profiles, err
}
