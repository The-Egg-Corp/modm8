package profile

import (
	"encoding/json"
	"fmt"
	"modm8/backend"
	"os"
	"path/filepath"
)

type ManifestOperation int

const (
	MANIFEST_OP_MOD_ADD ManifestOperation = iota
	MANIFEST_OP_MOD_REMOVE
)

// type ModStore map[string]IMod

// This struct is responsible for the management of profiles such as retreiving, creating, deleting and saving manifests.
//
// NOTE: All methods call regular functions so we can test said regular functions easily (since they aren't bound to the struct),
// and know that the methods will work exactly the same. As profiles are a core feature, we should take extra caution and ensure tests pass.
type ProfileManager struct{}

func NewManager() *ProfileManager {
	return &ProfileManager{}
}

func (pm *ProfileManager) GetProfiles(gameTitle string) (map[string]ProfileManifest, error) {
	return GetProfiles(gameTitle)
}

func (pm *ProfileManager) GetProfile(gameTitle, profileName string) (*ProfileManifest, error) {
	return GetManifest(gameTitle, profileName)
}

func (pm *ProfileManager) NewProfile(gameTitle, profileName string) error {
	return SaveManifest(gameTitle, profileName, NewProfileManifest())
}

func (pm *ProfileManager) SaveProfile(gameTitle, profileName string, prof ProfileManifest) error {
	return SaveManifest(gameTitle, profileName, prof)
}

func (pm *ProfileManager) DeleteProfile(gameTitle, profileName string) error {
	return DeleteProfile(gameTitle, profileName)
}

func (pm *ProfileManager) AddThunderstoreModToProfile(gameTitle string, profileName string, verFullName string) error {
	return UpdateThunderstoreProfileMods(MANIFEST_OP_MOD_ADD, gameTitle, profileName, verFullName)
}

func (pm *ProfileManager) AddNexusModToProfile(gameTitle string, profileName string, verFullName string) error {
	return UpdateNexusProfileMods(MANIFEST_OP_MOD_ADD, gameTitle, profileName, verFullName)
}

func (pm *ProfileManager) RemoveThunderstoreModFromProfile(gameTitle string, profileName string, verFullName string) error {
	return UpdateThunderstoreProfileMods(MANIFEST_OP_MOD_REMOVE, gameTitle, profileName, verFullName)
}

func (pm *ProfileManager) RemoveNexusModFromProfile(gameTitle string, profileName string, verFullName string) error {
	return UpdateNexusProfileMods(MANIFEST_OP_MOD_REMOVE, gameTitle, profileName, verFullName)
}

func UpdateThunderstoreProfileMods(op ManifestOperation, gameTitle string, profileName string, verFullName string) error {
	prof, err := GetManifest(gameTitle, profileName)
	if err != nil {
		return err
	}

	if op == MANIFEST_OP_MOD_ADD {
		prof.AddThunderstoreMod(verFullName)
	}

	if op == MANIFEST_OP_MOD_REMOVE {
		prof.RemoveThunderstoreMod(verFullName)
	}

	SaveManifest(gameTitle, profileName, *prof)
	return nil
}

func UpdateNexusProfileMods(op ManifestOperation, gameTitle string, profileName string, verFullName string) error {
	prof, err := GetManifest(gameTitle, profileName)
	if err != nil {
		return err
	}

	if op == MANIFEST_OP_MOD_ADD {
		prof.AddNexusMod(verFullName)
	}

	if op == MANIFEST_OP_MOD_REMOVE {
		prof.RemoveNexusMod(verFullName)
	}

	SaveManifest(gameTitle, profileName, *prof)
	return nil
}

func GameProfilesPath(gameTitle string) string {
	cacheDir, _ := os.UserConfigDir()
	path := filepath.Join(cacheDir, "modm8", "Games", gameTitle, "Profiles")

	return path
}

// Returns the full path to the manifest file for a given profile.
//
// For example, if gameTitle is "Skyrim" and profileName is "MyProfile", the output will be:
//
// <CONFIG_DIR>\modm8\Games\Skyrim\Profiles\MyProfile\profinfo.json
func PathToManifest(gameTitle, profileName string) string {
	return filepath.Join(GameProfilesPath(gameTitle), profileName, manifestName)
}

// Returns path info relating to the profile (directory and manifest name).
func ProfilePathInfo(gameTitle, profileName string) (string, string) {
	return filepath.Split(PathToManifest(gameTitle, profileName))
}

func GetProfileNames(gameTitle string) ([]string, error) {
	profilesPath := GameProfilesPath(gameTitle)

	// The user probably hasn't created any profiles yet.
	exists, _ := backend.ExistsAtPath(profilesPath)
	if !exists {
		return []string{}, nil
	}

	// Find all profile dirs containing a manifest file.
	paths, err := GetManifestDirs(profilesPath)
	if err != nil {
		return []string{}, err
	}

	return backend.GetBaseNames(paths), nil
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

func SaveManifest(gameTitle, profileName string, prof ProfileManifest) error {
	data, err := json.MarshalIndent(prof, "", "    ")
	if err != nil {
		return err
	}

	// Profiles dir, manifest name.
	dir, file := ProfilePathInfo(gameTitle, profileName)

	// Create the profiles dir (and its parents if missing) so we can save the manifest file inside it.
	err = backend.MkDirAll(dir)
	if err != nil {
		return err
	}

	return backend.WriteFile(filepath.Join(dir, file), data)
}

func DeleteProfile(gameTitle, profileName string) error {
	dir, _ := ProfilePathInfo(gameTitle, profileName)
	return os.RemoveAll(dir)
}

func GetManifest(gameTitle, profileName string) (*ProfileManifest, error) {
	contents, err := backend.ReadFile(PathToManifest(gameTitle, profileName))
	if err != nil {
		return nil, err
	}

	var manifest ProfileManifest
	if err := json.Unmarshal(contents, &manifest); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &manifest, nil
}

// Returns all of the directories at this path which contain a manifest.
func GetManifestDirs(path string) ([]string, error) {
	var dirs []string
	var root = filepath.Base(path)

	err := filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
		// Skip self and any files, we only want dirs.
		if !entry.IsDir() || entry.Name() == root {
			return nil
		}

		filePath := filepath.Join(path, manifestName)
		if _, err := os.Stat(filePath); err == nil {
			dirs = append(dirs, path)
		}

		return nil
	})

	return dirs, err
}
