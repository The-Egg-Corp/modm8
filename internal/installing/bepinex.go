package installing

import (
	"fmt"
	"modm8/internal/common/downloader"
	"modm8/internal/common/fileutil"
	"modm8/internal/common/paths"
	"os"
	"path/filepath"

	"github.com/cavaliergopher/grab/v3"
)

const BEPINEX_ROOT_NAME = "BepInEx"

const (
	INVALID    BepinexModLayout = iota
	STRUCTURED                  // Root dir (BEPINEX_ROOT_NAME) exists at some level. Sub dirs "plugins"
	FLAT                        // Root dir does not exist, but dlls are either top-level or in "plugins", "config" etc.
)

type BepinexModLayout int
type BepinexModInstaller struct {
}

func (ins *BepinexModInstaller) InstallLoader(downloadURL, dir string) (*grab.Response, error) {
	return InstallBepinexPack(downloadURL, dir)
}

// Downloads a zip at `downloadURL` and extracts it into a dir named `fullName` inside ModCache.
//
// This func subsequently calls Extract which will do the necessary IO operations
// and finally symlink the mod to `dir` (usually the profile directory).
func (ins *BepinexModInstaller) InstallMod(downloadURL, fullName, dir string) (*grab.Response, error) {
	// The path to mod in the cache dir. Ex: "../modm8/ModCache/Owen3H-IntroTweaks-1.5.0"
	modPath := filepath.Join(paths.ModCacheDir(), fullName)
	if exists, _ := fileutil.ExistsAtPath(modPath); exists {
		return nil, fmt.Errorf("mod '%s' already installed in %s", fullName, modPath)
	}

	// Download and unzip as is, then delete the zip.
	resp, err := downloader.DownloadAndUnzip(downloadURL, modPath, true)
	if err != nil {
		return resp, err
	}

	layout, err := DetectModLayout(modPath)
	if err != nil {
		return resp, err
	}

	switch layout {
	case FLAT:
		InstallFlatMod(fullName, dir, modPath)
	case STRUCTURED:

	}

	return resp, nil
}

func (ins *BepinexModInstaller) UninstallMod(fullName, dir string) error {
	// Drop symlink/junction

	return nil
}

func DetectModLayout(path string) (BepinexModLayout, error) {
	fi, err := os.Lstat(filepath.Join(path, BEPINEX_ROOT_NAME))
	if err != nil {
		return INVALID, err
	}

	if fi == nil {
		return INVALID, fmt.Errorf("BepInEx root does not exist within mod at %s", path)
	}
	if fi.IsDir() {
		return STRUCTURED, nil
	}

	return FLAT, err
}

// Extracts a mod zip with the assumption all dlls, configs etc. exist at the top-level without any Bepinex root.
func InstallFlatMod(fullName, dir, modPath string) error {
	// Check if "config", "plugins" etc. exist - deal with them appropriately.

	// Symlink "plugins"
	err := fileutil.LinkDir(filepath.Join(dir, fullName), modPath)
	if err != nil {
		return err
	}

	// Copy "plugins" to profile

	return nil
}

// Extracts a mod zip with the assumption that it has a valid structure (dlls, configs etc. nested within Bepinex root).
func InstallStructuredMod(rootPath string) error {
	return nil
}

// Installs BepInEx's own loader package at `path`, which is usually points to a profile dir.
func InstallBepinexPack(downloadURL, path string) (*grab.Response, error) {
	resp, err := downloader.DownloadAndUnzip(downloadURL, path, true)
	if err != nil {
		return resp, err
	}

	// BepInExPack should exist in what we just unzipped.
	bepinexPackDir := filepath.Join(path, "BepInExPack")

	// All files/dirs within the pack.
	entries, err := os.ReadDir(bepinexPackDir)
	if err != nil {
		return resp, err
	}

	// We now have a setup dir with BepInExPack inside of it.
	// The contents of BepInExPack need to go up 2 into the base profile dir.
	for _, entry := range entries {
		srcPath := filepath.Join(bepinexPackDir, entry.Name())
		dstPath := filepath.Join(path, entry.Name())

		err = os.Rename(srcPath, dstPath)
		if err != nil {
			return resp, err
		}
	}

	fileutil.MkDir(filepath.Join(path, BEPINEX_ROOT_NAME, "plugins"))

	// Then the leftover original bepinex setup dir can now be deleted.
	err = os.RemoveAll(bepinexPackDir)
	if err != nil {
		return resp, err
	}

	return resp, err
}
