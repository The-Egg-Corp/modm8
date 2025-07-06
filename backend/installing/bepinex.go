package installing

import (
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"modm8/backend/loaders"
	"os"
	"path/filepath"

	"github.com/cavaliergopher/grab/v3"
)

type BepinexModInstaller struct {
}

const BEPINEX_ZIP_OUTPUT_NAME = "BepInEx-Setup"

func (ins *BepinexModInstaller) Extract() error {
	return nil
}

func (ins *BepinexModInstaller) Install(downloadURL, dir, fullName string) error {
	// Install own loader package
	if loaders.IsLoaderPackage(loaders.BEPINEX, fullName) {
		_, err := InstallBepinexPack(downloadURL, dir)
		return err
	}

	// Install regular bepinex mod

	return nil
}

func (ins *BepinexModInstaller) Uninstall() error {
	return nil
}

func InstallBepinexPack(downloadURL, dir string) (*grab.Response, error) {
	outputPath := filepath.Join(dir, BEPINEX_ZIP_OUTPUT_NAME)

	res, err := downloader.DownloadUnzipDelete(outputPath, downloadURL)
	if err != nil {
		return res, err
	}

	bepinexPackDir := filepath.Join(outputPath, "BepInExPack")
	fileutil.MkDir(filepath.Join(bepinexPackDir, "plugins"))

	// We now have a dir with BepInExPack inside of it.
	// The contents of that dir need to go up 2 into the base profile dir.
	entries, err := os.ReadDir(bepinexPackDir)
	if err != nil {
		return res, err
	}

	// Take everything in BepInExPack and move it up two into `dir`.
	for _, entry := range entries {
		srcPath := filepath.Join(bepinexPackDir, entry.Name())
		dstPath := filepath.Join(dir, entry.Name())

		err = os.Rename(srcPath, dstPath)
		if err != nil {
			return res, err
		}
	}

	// Then the leftover original bepinex setup dir can now be deleted.
	err = os.RemoveAll(outputPath)
	if err != nil {
		return res, err
	}

	return res, err
}
