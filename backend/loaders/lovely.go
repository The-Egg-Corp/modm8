package loaders

import "fmt"

type LovelyLoaderInstructions struct {
}

// TODO: Implement this
func (instr LovelyLoaderInstructions) Generate(profileDir string) (*LoaderInstructions, error) {
	return &LoaderInstructions{}, fmt.Errorf("instructions for loader LOVELY not yet implemented")
}
