package loaders

import (
	"fmt"
	"path/filepath"
)

type LovelyLoader struct {
}

// TODO: Implement this
func (ldr LovelyLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}

func (ldr LovelyLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join("lovely", "mods")
}
