package loaders

import (
	"fmt"
	"path/filepath"
)

type MelonLoader struct {
}

// TODO: Implement this
func (ldr MelonLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader MELON not yet implemented")
}

func (ldr MelonLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join(profileDir, "Mods")
}
