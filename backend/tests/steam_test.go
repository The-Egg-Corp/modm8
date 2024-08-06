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
		t.Fatalf("error getting steam directory: %v", err)
	}

	if path == nil {
		t.Error("Could not find directory, please specify one manually.")
	} else {
		fmt.Println("Steam installed at:", *path)
	}
}

func TestLaunchSteamGame(t *testing.T) {
	cmd, err := runner.LaunchSteamGame(1966720, []string{"--doorstop-enable", "false"})
	if err != nil {
		t.Fatalf("error launching game: %v", err)
		return
	}

	fmt.Println("Executing command:", cmd.String())

	rootPid := cmd.Process.Pid
	fmt.Println("Root PID:", rootPid)
}
