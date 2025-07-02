package game

import (
	"fmt"
	"modm8/backend/common/fileutil"
	"path/filepath"
	"strings"
)

type PlatformArch string

const (
	X64 PlatformArch = "x64"
	X86 PlatformArch = "x86"
)

type BepinexConfig struct {
	RootComments []string                      `json:"root_comments" mapstructure:"root_comments"`
	Entries      map[string]BepinexConfigEntry `json:"entries" mapstructure:"entries"`
}

type BepinexConfigEntry struct {
	Section          string   `json:"section" mapstructure:"section"`
	Value            string   `json:"value" mapstructure:"value"`
	DefaultValue     string   `json:"default_value" mapstructure:"default_value"`
	Comments         []string `json:"comments" mapstructure:"comments"`
	AcceptableValues []string `json:"acceptable_values" mapstructure:"acceptable_values"`
}

func ParseBepinexConfig(path string) (*BepinexConfig, error) {
	contents, err := fileutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	entries := make(map[string]BepinexConfigEntry)

	// Distinct key to avoid conflicting with a possible section.
	currentSection := "__root"
	currentDefaultValue := ""

	var rootComments []string
	var comments []string
	var acceptableValues []string

	reset := func() {
		comments = nil
		acceptableValues = make([]string, 0)
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			if len(comments) < 1 {
				continue
			}

			if currentSection == "__root" {
				rootComments = append(rootComments, comments...)
			}

			reset()
		}

		//#region Parse comments
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			comments = append(comments, line)

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
			if len(comments) > 0 && currentSection == "__root" {
				rootComments = append(rootComments, comments...)

				reset()
			}

			// Grab everything in between the brackets.
			currentSection = line[1 : len(line)-1]
			continue
		}
		//#endregion

		//#region Parse option (key/value pair)
		if strings.Index(line, "=") < 1 {
			continue // Ignore malformed line.
		}

		if key, value, found := strings.Cut(line, "="); found {
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)

			// Put the section within the key to avoid map nesting.
			flatKey := fmt.Sprintf("%s.%s", currentSection, key)

			entries[flatKey] = BepinexConfigEntry{
				Value:            value,
				DefaultValue:     currentDefaultValue,
				Comments:         comments,
				AcceptableValues: acceptableValues,
			}

			reset()
		}
		//#endregion
	}

	return &BepinexConfig{
		RootComments: rootComments,
		Entries:      entries,
	}, nil
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
		if exists, _ := fileutil.ExistsAtPath(path); !exists {
			missing = append(missing, fileName)
		}
	}

	// If no missing files, BepInEx is likely installed.
	return len(missing) == 0, missing
}
