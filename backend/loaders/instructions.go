package loaders

import "fmt"

type LoaderInstructions struct {
	ModdedParams  string
	VanillaParams string
}

type ILoaderInstructions interface {
	Generate(profileDir string) LoaderInstructions
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

	instructions := l.Generate(profileDir)
	return &instructions, nil
}

func GenDoorstopV3() LoaderInstructions {
	return LoaderInstructions{
		ModdedParams:  "--doorstop-enable true --doorstop-target BepInEx\\core\\BepInEx.Preloader.dll",
		VanillaParams: "--doorstop-enable false",
	}
}

func GenDoorstopV4() LoaderInstructions {
	return LoaderInstructions{
		ModdedParams:  "--doorstop-enabled true --doorstop-target-assembly BepInEx\\core\\BepInEx.Preloader.dll",
		VanillaParams: "--doorstop-enabled false",
	}
}
