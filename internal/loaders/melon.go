package loaders

import (
	"fmt"
	"path/filepath"
)

type MelonModLoader struct {
}

// func (ldr MelonModLoader) IsLoaderPackage(fullName string) bool {
// 	return fullName == "LavaGang-MelonLoader"
// }

func (ldr MelonModLoader) GetModLinkPath(profileDir string) string {
	return filepath.Join(profileDir, "Mods")
}

// TODO: Implement this
func (ldr MelonModLoader) GenerateInstructions(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader MELON not yet implemented")
}
