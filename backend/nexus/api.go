package nexus

import (
	"bufio"
	"context"
	"fmt"
	"modm8/backend/app"
	"os"
	"strings"

	v1 "github.com/the-egg-corp/gonexus/v1"
)

type API struct {
	Ctx    context.Context
	Cache  map[string]any
	Client v1.NexusClient
}

func NewCache() map[string]any {
	return make(map[string]any, 0)
}

func NewAPI(ctx context.Context) (*API, error) {
	//fmt.Print("Gathering nexus key...\n")

	// TODO: A single file for a key is a bit odd. Store it in a better way. Win/Linux registry maybe?

	// Grab nexus key from `<configDir>/modm8/nex.key` file.
	key, err := ScanFileForValidLine(app.KeyPath())
	if err != nil {
		return nil, err
	}

	client, err := v1.NewNexusClient(key)
	if err != nil {
		return nil, err
	}

	return &API{
		Ctx:    ctx,
		Cache:  NewCache(),
		Client: *client,
	}, nil
}

// Scans the contents of the file at the given path line by line. At the first instance of a non-blank line,
// the line is returned with no line endings or carriage return, making this function platform-independent.
func ScanFileForValidLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)

		if line != "" {
			return line, nil // Scan should have already removed \n and \r.
		}
	}

	if err := s.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("%s: file exists, but found no valid line found inside", path)
}
