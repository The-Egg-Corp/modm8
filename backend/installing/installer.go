package installing

import (
	"fmt"
	"modm8/backend/loaders"
)

type IModInstaller interface {
	Extract() error
	Install(dir, fullName, downloadURL string) error
	Uninstall() error
}

var MOD_INSTALLERS = map[loaders.ModLoaderType]IModInstaller{
	loaders.BEPINEX: &BepinexModInstaller{},
}

func GetModInstaller(loader loaders.ModLoaderType) (IModInstaller, error) {
	ins, ok := MOD_INSTALLERS[loader]
	if !ok {
		return nil, fmt.Errorf("failed to get mod installer. no loader exists with index %d", loader.Index())
	}

	return ins, nil
}
