package loaders

import "fmt"

type LoaderInstructions struct {
	ModdedParams  []string
	VanillaParams []string
}

type ILoader interface {
	GenerateInstructions(profileDir string) (*LoaderInstructions, error)
	GetModLinkPath(profileDir string) string
}

var MOD_LOADERS = map[ModLoader]ILoader{
	BEPINEX: &BepinexLoader{},
	MELON:   &MelonLoader{},
	LOVELY:  &LovelyLoader{},
}

func GetLoaderInstructions(loader ModLoader, profileDir string) (*LoaderInstructions, error) {
	l, ok := MOD_LOADERS[loader]
	if !ok {
		return nil, fmt.Errorf("failed to get instructions. invalid loader with index %d", loader.Index())
	}

	instructions, err := l.GenerateInstructions(profileDir)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}

func GetModLinkPath(loader ModLoader, profileDir string) (string, error) {
	l, ok := MOD_LOADERS[loader]
	if !ok {
		return "", fmt.Errorf("failed to get instructions. invalid loader with index %d", loader.Index())
	}

	return l.GetModLinkPath(profileDir), nil
}
