package installing

import (
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"os"
	"path/filepath"

	"github.com/cavaliergopher/grab/v3"
)

type BepinexModInstaller struct {
}

const BEPINEX_ZIP_OUTPUT_NAME = "BepInEx-Setup"
const BEPINEX_ROOT_NAME = "BepInEx"

func (ins *BepinexModInstaller) Install(downloadURL, fullName, cacheDir string) (*grab.Response, error) {
	// Install own loader package.
	// if loaders.IsLoaderPackage(loaders.BEPINEX, fullName) {
	// 	profDir := profile.PathToProfile(gameTitle, profName)
	// 	_, err := InstallBepinexPack(downloadURL, profDir)
	// 	return err
	// }

	// Download zip and extract contents in a new mod dir.
	path := filepath.Join(cacheDir, fullName)
	res, err := downloader.DownloadAndUnzip(downloadURL, path, true)
	if err != nil {
		return res, err
	}

	// If flat, we don't need to do anything more, just link.
	// Otherwise, walk mod for dlls, configs and normalize to valid structure.

	return res, nil
}

func (ins *BepinexModInstaller) Uninstall(fullName, dir string) error {
	// Drop symlink/junction

	return nil
}

// Installs BepInEx's own loader package at `path`, which is usually points to a profile dir.
func InstallBepinexPack(downloadURL, path string) (*grab.Response, error) {
	res, err := downloader.DownloadAndUnzip(downloadURL, path, true)
	if err != nil {
		return res, err
	}

	// BepInExPack should exist in what we just unzipped.
	bepinexPackDir := filepath.Join(path, "BepInExPack")

	// All files/dirs within the pack.
	entries, err := os.ReadDir(bepinexPackDir)
	if err != nil {
		return res, err
	}

	// We now have a setup dir with BepInExPack inside of it.
	// The contents of BepInExPack need to go up 2 into the base profile dir.
	for _, entry := range entries {
		srcPath := filepath.Join(bepinexPackDir, entry.Name())
		dstPath := filepath.Join(path, entry.Name())

		err = os.Rename(srcPath, dstPath)
		if err != nil {
			return res, err
		}
	}

	fileutil.MkDir(filepath.Join(path, BEPINEX_ROOT_NAME, "plugins"))

	// Then the leftover original bepinex setup dir can now be deleted.
	err = os.RemoveAll(bepinexPackDir)
	if err != nil {
		return res, err
	}

	return res, err
}

// Installs a mod with the assumption all dlls, configs etc. exist at the top-level.
func InstallFlat() error {
	return nil
}

// Installs a mod with the assumption that it has a valid structure (dlls, configs etc. nested within BepInEx dir).
func InstallStructured(rootPath string) error {
	return nil
}
