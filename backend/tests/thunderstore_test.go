package backend

import (
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"modm8/backend/game"
	"modm8/backend/thunderstore"
	"path/filepath"
	"testing"
	"time"

	v1 "github.com/the-egg-corp/thundergo/v1"
)

const tsDownloadDomain = "https://thunderstore.io/package/download/"
const testPkg1 = "sfDesat-Orion-2.1.4"

var comm = v1.Community{
	Identifier: "lethal-company",
}

func TestInstallWithDependencies(t *testing.T) {
	pkgs, err := comm.AllPackages()
	if err != nil {
		t.Fatalf("error getting packages in community:\n%v", err)
	}

	// Test dependencies by installing a modpack
	pkg := pkgs.Get("Megalophobia", "MEGALOPHOBIA")
	if pkg == nil {
		t.Fatal("could not find package in community")
	}

	// isModpack := pkg.IsModpack(false)
	// if !isModpack {
	// 	t.Fatal("specified package is not a modpack")
	// }

	t.Logf("\nFound package: %s\nDownloading dependencies...\n\n", pkg.LatestVersion().FullName)

	var errs []error
	var downloadCount int

	startTime := time.Now()

	thunderstore.CurModCacheDir = game.ModCacheDir("Lethal Company")
	thunderstore.InstallWithDependencies(pkg.LatestVersion(), pkgs, &errs, &downloadCount)

	t.Logf("\nDownloaded %v packages in %v\n", downloadCount, time.Since(startTime))
}

func TestUnzipAndDelete(t *testing.T) {
	path := filepath.Join(game.ModCacheDir("Lethal Company"), testPkg1)
	zipPath := path + downloader.CUSTOM_ZIP_EXT

	err := fileutil.UnzipAndDelete(zipPath, path)
	if err != nil {
		t.Fatalf("\nerror unpacking zip:\n\n%v", err)
	}

	t.Logf("\n\nSuccessfully unpacked and deleted:\n  %s", zipPath)
}
