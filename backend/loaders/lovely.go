package loaders

import (
	"fmt"
	"path/filepath"
)

type LovelyLoader struct {
}

// TODO: Implement this
func (instr LovelyLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}

func (instr LovelyLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join("lovely", "mods")
}
