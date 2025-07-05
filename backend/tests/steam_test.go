package backend

import (
	"fmt"
	"modm8/backend/common/profile"
	"modm8/backend/loaders"
	steam "modm8/backend/runners/steam"
	"strings"
	"testing"
)

// Risk of Rain 2: 632360
// Lethal Company: 1966720
const testGame = 1966720

func TestGetInstallDirectory(t *testing.T) {
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

func TestLaunchGameWindows(t *testing.T) {
	steamPath, err := steam.GetInstallDirectory()
	if err != nil {
		t.Fatalf("error launching game. failed to find path to Steam: %v", err)
	}

	instructions, err := loaders.GetLoaderInstructions(loaders.BEPINEX, profile.PathToProfile("Lethal Company", "test"))
	if err != nil {
		t.Fatalf("error launching game:\n%v", err)
		return
	}

	cmd, err := steam.LaunchGame(steamPath, "steam.exe", testGame, instructions.ModdedParams)
	if err != nil {
		t.Fatalf("error launching game:\n%v", err)
		return
	}

	fullCmd := fmt.Sprint(cmd.Name, " ", strings.Join(cmd.Args, " "))
	fmt.Printf("[PID: %d] Executing '%s'\n", cmd.Status().PID, fullCmd)
}

func TestGetLibs(t *testing.T) {
	libPaths, err := steam.GetSteamLibPaths()
	if err != nil {
		t.Fatalf("error getting steam libs:\n%v", err)
	}

	fmt.Print("steam lib paths:\n")
	for _, path := range libPaths {
		fmt.Printf("%s\n", path)
	}
}
