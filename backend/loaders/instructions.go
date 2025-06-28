package loaders

import "fmt"

type LoaderInstructions struct {
	ModdedParams  []string
	VanillaParams []string
}

type ILoaderInstructions interface {
	Generate(profileDir string) (*LoaderInstructions, error)
}

var LOADER_INSTRUCTIONS = map[ModLoader]ILoaderInstructions{
	BEPINEX: &BepinexLoaderInstructions{},
	MELON:   &MelonLoaderInstructions{},
	LOVELY:  &LovelyLoaderInstructions{},
}

func GetLoaderInstructions(loader ModLoader, profileDir string) (*LoaderInstructions, error) {
	l, ok := LOADER_INSTRUCTIONS[loader]
	if !ok {
		return nil, fmt.Errorf("failed to get instructions. invalid loader with index %d", loader.Index())
	}

	instructions, err := l.Generate(profileDir)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
