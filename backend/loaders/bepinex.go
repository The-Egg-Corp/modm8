package loaders

import "modm8/backend/utils"

type BepinexLoaderInstructions struct {
}

func (instr BepinexLoaderInstructions) Generate(profileDir string) LoaderInstructions {
	doorstopVer, _ := utils.GetUnityDoorstopVersion(profileDir)
	if doorstopVer == 4 {
		return GenDoorstopV4()
	}

	return GenDoorstopV3()
}
