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

func (ins *BepinexModInstaller) Extract() error {
	return nil
}

func (ins *BepinexModInstaller) Install() error {
	return nil
}

func (ins *BepinexModInstaller) Uninstall() error {
	return nil
}

const BEPINEX_PACK_OUTPUT_NAME = "BepInEx-Setup"

func InstallBepinexPack(downloadURL, dir string) (*grab.Response, error) {
	// Send download request to `downloadURL` and save the zip as `outputFileName` in `dir`.
	resp, err := downloader.DownloadZip(downloadURL, filepath.Join(dir, BEPINEX_PACK_OUTPUT_NAME))
	if err != nil {
		return resp, err
	}
	if err = resp.Err(); err != nil {
		return resp, err
	}

	// Unzip the downloaded file into the same directory.
	outputPath := filepath.Join(dir, BEPINEX_PACK_OUTPUT_NAME)
	err = fileutil.UnzipAndDelete(outputPath+downloader.CUSTOM_ZIP_EXT, outputPath)
	if err != nil {
		return resp, err
	}

	// Path to the BepInExPack folder inside the extracted dir.
	bepinexPackDir := filepath.Join(outputPath, "BepInExPack")

	// Create the plugins folder if it didn't exist in the pack.
	fileutil.MkDir(filepath.Join(bepinexPackDir, "plugins"))

	// We now have a dir with BepInExPack inside of it.
	// The contents of that dir need to go up 2 into the base profile dir.
	entries, err := os.ReadDir(bepinexPackDir)
	if err != nil {
		return resp, err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(bepinexPackDir, entry.Name())
		dstPath := filepath.Join(dir, entry.Name())

		err = os.Rename(srcPath, dstPath)
		if err != nil {
			return resp, err
		}
	}

	// Then the leftover original bepinex pack dir can now be deleted.
	err = os.RemoveAll(outputPath)
	if err != nil {
		return resp, err
	}

	return resp, err
}
