package backend

import (
	"fmt"
	"modm8/backend"
	"testing"
)

func TestSteamDir(t *testing.T) {
	path, err := backend.FindSteamDirectory()
	if err != nil {
		t.Error("Could not find directory, please specify one manually.")
	}

	fmt.Println("Steam installed at:", *path)
}
