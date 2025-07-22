package installing

import (
	"fmt"
	"modm8/internal/loaders"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

type PackageInstallMeta struct {
	Loader       loaders.ModLoaderType `json:"platform"`     // The loader this package uses to be setup and function correctly.
	FullName     string                `json:"fullName"`     // The full name of this package. Ex: "Owen3H-IntroTweaks" or "Owen3H-IntroTweaks-1.5.0"
	Dependencies []string              `json:"dependencies"` // Contains the names of the packages that this package requires.
	DownloadURL  string                `json:"downloadURL"`  // The URL where this package lives and can be retrieved.
}

func (meta *PackageInstallMeta) HasVersionSuffix() bool {
	return len(strings.Split(meta.FullName, "-")) > 1
}

type IModInstaller interface {
	InstallLoader(downloadURL, dir string) (*grab.Response, error)
	InstallMod(downloadURL, fullName, dir string) (*grab.Response, error)
	UninstallMod(fullName, dir string) error
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
