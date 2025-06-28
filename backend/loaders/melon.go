package loaders

import (
	"fmt"
	"path/filepath"
)

type MelonLoaderInstructions struct {
}

// TODO: Implement this
func (instr MelonLoaderInstructions) Generate(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader MELON not yet implemented")
}

func (instr MelonLoaderInstructions) GetModLinkPath(profileDir string) string {
	return filepath.Join(profileDir, "Mods")
}
