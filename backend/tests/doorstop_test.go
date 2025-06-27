package backend

import (
	"fmt"
	"modm8/backend/loaders"
	"os"
	"path/filepath"
	"testing"
)

func TestGetUnityDoorstopVersion(t *testing.T) {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("failed test setup: %v\n", err)
	}

	// TODO: Currently tests r2. We need to change this to modm8 when BepInEx-Pack setup is implemented.
	path := filepath.Join(cfgPath, "r2modmanPlus-local", "RiskOfRain2", "profiles", "Default")

	ver, err := loaders.GetUnityDoorstopVersion(path)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Unity Doorstop version: %d\n", ver)
}
