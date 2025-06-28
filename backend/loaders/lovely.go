package loaders

import (
	"fmt"
	"path/filepath"
)

type LovelyLoaderInstructions struct {
}

// TODO: Implement this
func (instr LovelyLoaderInstructions) Generate(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}

func (instr LovelyLoaderInstructions) GetModLinkPath(profileDir string) string {
	return filepath.Join("lovely", "mods")
}
