package backend

import (
	"fmt"
	"modm8/backend"
	"testing"
)

var runner = backend.NewSteamRunner()

func TestSteamDir(t *testing.T) {
	path, err := backend.GetSteamDirectory()
	if err != nil {
		t.Error("Could not find directory, please specify one manually.")
	}

	fmt.Println("Steam installed at:", *path)
}

func TestLaunchSteamGame(t *testing.T) {
	err := runner.LaunchSteamGame(1966720)
	if err != nil {
		t.Fatalf("error launching game: %v", err)
	}
}
