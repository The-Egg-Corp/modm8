package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DOORSTOP_VER_FILE_NAME = ".doorstop_version"

// Reads the ".doorstop_version" file in dirPath (usually the game or profile directory) and returns the major version specified within it.
// This file only exists from Unity Doorstop v4 onwards, and should contain a single line representing a semantic version such as "v4.0.0".
//
// Even if the file content is malformed with spaces or other characters before or after, this func will try to find the first line with a valid semantic version.
// If we dont find it or the version is somehow <= 3, defaults and returns 3.
func GetUnityDoorstopVersion(dirPath string) (int, error) {
	semverLine, err := FindFirstSemverLine(dirPath)
	if err != nil {
		return 3, err
	}

	majorVer, err := strconv.Atoi(*semverLine)
	if err != nil || majorVer <= 3 {
		return 3, err
	}

	return majorVer, nil
}

func FindFirstSemverLine(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineContent := strings.TrimSpace(scanner.Text())
		if lineContent == "" {
			continue
		}

		semverParts := strings.Split(lineContent, ".")
		if len(semverParts) != 3 {
			continue
		}

		if _, err1 := strconv.Atoi(semverParts[0]); err1 != nil {
			continue
		}
		if _, err2 := strconv.Atoi(semverParts[1]); err2 != nil {
			continue
		}
		if _, err3 := strconv.Atoi(semverParts[2]); err3 != nil {
			continue
		}

		return &lineContent, nil
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file at %s\n%v", path, err)
	}

	return nil, errors.New("no semantic version line found")
}
