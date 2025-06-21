package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DOORSTOP_VER_FILE_NAME = ".doorstop_version"

// Reads the ".doorstop_version" file in dirPath.
//
// This file only exists on v4 and above.
// If we dont find it or the version is somehow <= 3, just default to 3.
func GetUnityDoorstopVersion(dirPath string) (int, error) {
	data, err := os.ReadFile(filepath.Join(dirPath, DOORSTOP_VER_FILE_NAME))
	if err != nil {
		return 3, err
	}

	// Clean up content. Trim whitespace and other space-like characters.
	content := strings.TrimSpace(string(data))

	// Split the semantic version string and only grab major/first.
	parts := strings.SplitN(content, ".", 2)
	majorStr := strings.TrimSpace(parts[0])

	majorVer, err := strconv.Atoi(majorStr)
	if err != nil || majorVer <= 3 {
		return 3, err
	}

	return majorVer, nil
}
