package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Scans the contents of the file at the given path line by line. At the first instance of a non-blank line,
// the line is returned with no line endings or carriage return, making this function platform-independent.
func FindFirstValidLine(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lineContent := strings.TrimSpace(s.Text())
		if lineContent != "" {
			return &lineContent, nil // Scan should have already removed \n and \r.
		}
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error reading file at %s\n%v", path, err)
	}

	return nil, fmt.Errorf("no valid line found in file: %s", path)
}

// Like [FindFirstValidLine], this function scans the contents of the file at the given path line by line,
// but also adds an extra constraint where the line must be in a proper semver format.
//
// In summary, the first line that is both non-blank and valid semver will be returned.
func FindFirstSemverLine(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lineContent := strings.TrimSpace(s.Text())
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

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error reading file at %s\n%v", path, err)
	}

	return nil, fmt.Errorf("no semver line found in file: %s", path)
}
