package backend

import (
	"fmt"
	steam "modm8/backend/runners/steam"
	"testing"
)

func TestSteamDir(t *testing.T) {
	path, err := steam.GetInstallDirectory()

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
	cmd, err := steam.LaunchGame(1966720, []string{"--doorstop-enable", "false"})
	if err != nil {
		t.Fatalf("error launching game: %v", err)
		return
	}

	fmt.Println("Executing command:", cmd.String())

	rootPid := cmd.Process.Pid
	fmt.Println("Root PID:", rootPid)
}
