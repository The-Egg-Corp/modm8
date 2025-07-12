package installing

import (
	"fmt"
	"modm8/backend/loaders"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

type PackageInstallMeta struct {
	// The loader this package uses to be setup and function correctly.
	Loader loaders.ModLoaderType `json:"platform"`
	// The full name of this package. Ex: "Owen3H-IntroTweaks" or "Owen3H-IntroTweaks-1.5.0"
	FullName string `json:"fullName"`
	// Contains the names of the packages that this package requires.
	Dependencies []string `json:"dependencies"`
	// The URL where this package lives and can be retrieved.
	DownloadURL string `json:"downloadURL"`
}

func (meta *PackageInstallMeta) HasVersionSuffix() bool {
	return len(strings.Split(meta.FullName, "-")) > 1
}

type IModInstaller interface {
	//InstallSelf(downloadURL, dir string) error
	Install(downloadURL, fullName, dir string) (*grab.Response, error)
	Uninstall(fullName, dir string) error
	//Extract() error
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
