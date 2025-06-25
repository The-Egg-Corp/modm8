package backend

import (
	"modm8/backend/thunderstore"
	"os"
	"testing"
)

var schema = thunderstore.NewThunderstoreSchema()

func TestGetEcosystem(t *testing.T) {
	ecosys, err := schema.GetEcosystem()
	if err != nil {
		t.Fatalf("GetEcosystem returned error:\n%v", err)
	}
	if ecosys == nil {
		t.Fatal("GetEcosystem returned nil ecosystem")
	}

	path, err := thunderstore.GetFallbackEcosystemPath()
	if err != nil {
		t.Fatalf("failed to get fallback path:\n%v", err)
	}

	if _, err := os.Stat(*path); err != nil {
		t.Fatalf("expected fallback file to exist at %s, but got error: %v", *path, err)
	}
}
