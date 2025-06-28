package loaders

import (
	"fmt"
	"path/filepath"
)

type MelonLoader struct {
}

// TODO: Implement this
func (instr MelonLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader MELON not yet implemented")
}

func (instr MelonLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join(profileDir, "Mods")
}
