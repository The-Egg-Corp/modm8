package loaders

import (
	"fmt"
	"path/filepath"
)

type LovelyModLoader struct {
}

// func (ldr LovelyModLoader) IsLoaderPackage(fullName string) bool {
// 	return fullName == "Thunderstore-lovely"
// }

func (ldr LovelyModLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join("lovely", "mods")
}

// TODO: Implement this
func (ldr LovelyModLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}
