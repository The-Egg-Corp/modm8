package backend

import (
	"fmt"
	epic "modm8/backend/runners/epic"
	"strings"
	"testing"
)

func TestLaunchEpicGame(t *testing.T) {
	cmd, err := epic.LaunchGame("jaguar")
	if err != nil {
		t.Fatalf("error launching game: %v", err)
		return
	}

	fullCmd := fmt.Sprint(cmd.Name, " ", strings.Join(cmd.Args, " "))
	fmt.Printf("[PID: %d] Executing '%s'\n", cmd.Status().PID, fullCmd)
}