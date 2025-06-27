package loaders

import "fmt"

type MelonLoaderInstructions struct {
}

// TODO: Implement this
func (instr MelonLoaderInstructions) Generate(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader MELON not yet implemented")
}
