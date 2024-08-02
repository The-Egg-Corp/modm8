package backend

import (
	"fmt"
	"modm8/backend/core"
	"testing"
)

func TestLoad(t *testing.T) {
	settings := core.NewSettings()

	if err := settings.Load(); err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	expectedSettings := core.NewSettings()

	// Check loaded settings match expected defaults
	if *settings != *expectedSettings {
		t.Errorf("\nLoaded settings do not match expected default settings.\n\nGot: %v\nExpected: %v", settings, expectedSettings)
		return
	}

	fmt.Println(settings)
}
