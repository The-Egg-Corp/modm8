package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
