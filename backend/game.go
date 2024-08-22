package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

type BepinexConfig struct {
	RootComments []string                    `json:"root_comments" mapstructure:"root_comments"`
	Variables    map[string]BepinexConfigVar `json:"variables" mapstructure:"variables"`
}

type BepinexConfigVar struct {
	Section          string   `json:"section" mapstructure:"section"`
	Value            string   `json:"value" mapstructure:"value"`
	DefaultValue     string   `json:"default_value" mapstructure:"default_value"`
	Comments         []string `json:"comments" mapstructure:"comments"`
	AcceptableValues []string `json:"acceptable_values" mapstructure:"acceptable_values"`
}

type GameManager struct{}

func NewGameManager() *GameManager {
	return &GameManager{}
}

func (gm *GameManager) GameInstalled(dirPath string, exeKeywords []string) bool {
	var installed = false

	for _, keyword := range exeKeywords {
		path := filepath.Join(dirPath, keyword)
		if exists, _ := ExistsAtPath(path); exists {
			installed = true
			break
		}
	}

	return installed
}

func (gm *GameManager) ParseBepinexConfig(path string) (*BepinexConfig, error) {
	return ParseBepinexConfig(path)
}

func ParseBepinexConfig(path string) (*BepinexConfig, error) {
	contents, err := ReadFile(path)
	if err != nil {
		return nil, err
	}

	vars := make(map[string]BepinexConfigVar)

	// Distinct key to avoid conflicting with a possible section.
	currentSection := "__root"
	currentDefaultValue := ""

	var rootComments = make([]string, 0)
	var comments []string
	var acceptableValues []string

	lines := strings.Split(*contents, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Don't need to do anything with an empty line
		if len(line) == 0 {
			continue
		}

		//#region Parse comments
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			if currentSection != "__root" {
				comments = append(comments, line)
			} else {
				rootComments = append(comments, line)
			}

			if strings.Contains(line, "Setting type: Boolean") {
				acceptableValues = []string{"true", "false"}
			}

			if strings.Contains(line, "Default value: ") {
				currentDefaultValue = strings.TrimSpace(line[len("# Default value:"):])
			}

			continue
		}
		//#endregion

		//#region Parse section/category
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// Grab everything in between the brackets
			currentSection = line[1 : len(line)-1]
			continue
		}
		//#endregion

		//#region Parse option (key/value pair)
		if strings.Index(line, "=") < 1 {
			continue // Ignore malformed line
		}

		if key, value, found := strings.Cut(line, "="); found {
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)

			// Put the section within the key to avoid nested maps. Ex: parsed[section][key]
			flatKey := fmt.Sprintf("%s.%s", currentSection, key)

			vars[flatKey] = BepinexConfigVar{
				Section:          currentSection,
				Value:            value,
				DefaultValue:     currentDefaultValue,
				Comments:         comments,
				AcceptableValues: acceptableValues,
			}

			// Clear comments and acceptable values for the next entry
			comments = nil
			acceptableValues = make([]string, 0)
		}
		//#endregion
	}

	return &BepinexConfig{
		RootComments: rootComments,
		Variables:    vars,
	}, nil
}

func (gm *GameManager) BepinexConfigFiles(dirs []string) ([]string, error) {
	return WalkDirExt(filepath.Join(dirs...), []string{"cfg"})
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

	for _, fileName := range required {
		// Use abs path to get path to current required file.
		path := filepath.Join(absPath, fileName)

		// Check if file exists, add to missing slice if not.
		if exists, _ := ExistsAtPath(path); !exists {
			missing = append(missing, fileName)
		}
	}

	// If no missing files, BepInEx is likely installed.
	return len(missing) == 0, missing
}

func ExistsAtPath(absPath string) (bool, error) {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}
