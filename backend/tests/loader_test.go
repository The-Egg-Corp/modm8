package backend

import (
	"fmt"
	"modm8/backend/common/profile"
	"modm8/backend/loaders"
	"testing"
)

const testLoader = loaders.BEPINEX

func TestGetLoaderInstructions(t *testing.T) {
	instructions, err := loaders.GetLoaderInstructions(
		testLoader,
		profile.PathToProfile("Lethal Company", "test"),
	)

	if err != nil {
		t.Fatal(err)
	}

	if instructions == nil {
		t.Fatalf("instructions for loader %s returned nil", testLoader.Name())
	}

	fmt.Printf("Params for loader %s\n", testLoader.Name())
	fmt.Printf("Modded: %s\n", instructions.ModdedParams)
	fmt.Printf("Vanilla: %s\n\n", instructions.VanillaParams)
}
