package game

type GameInstruction struct {
	ModdedParams  string
	VanillaParams string
}

type GameInstructions interface {
	Generate() GameInstruction
}

type BepInExInstructions struct {
}

func (instr BepInExInstructions) Generate() GameInstruction {
	return GameInstruction{
		ModdedParams:  "--doorstop-enable true --doorstop-target BepInEx\\core\\BepInEx.Preloader.dll",
		VanillaParams: "--doorstop-enable false",
	}
}

type MelonLoaderInstructions struct {
}

// TODO: Implement this
func (instr MelonLoaderInstructions) Generate() GameInstruction {
	return GameInstruction{}
}
