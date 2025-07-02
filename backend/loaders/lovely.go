package loaders

import (
	"fmt"
	"path/filepath"
)

type LovelyModLoader struct {
}

// TODO: Implement this
func (ldr LovelyModLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}

func (ldr LovelyModLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join("lovely", "mods")
}
