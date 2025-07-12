package backend

import (
	"modm8/backend/app"
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"modm8/backend/installing"
	"modm8/backend/loaders"
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
	ecosys, err := thunderstore.GetSchema().GetEcosystem()
	if err != nil {
		t.Fatalf("failed to get ecosystem:\n%v", err)
	}

	r2mapping := ecosys.Games[comm.Identifier].R2Modman[0]
	loader := loaders.GetModLoaderType(r2mapping.PackageLoader)

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

	ver := pkg.LatestVersion()
	t.Logf("\nFound package: %s\nDownloading dependencies...\n\n", ver.FullName)

	var errs []error
	var downloadCount int

	meta := installing.PackageInstallMeta{
		Loader:       loader,
		FullName:     ver.FullName,
		Dependencies: ver.Dependencies,
		DownloadURL:  ver.DownloadURL,
	}

	startTime := time.Now()
	thunderstore.InstallWithDependencies(meta, pkgs, &errs, &downloadCount)

	t.Logf("\nDownloaded %v packages in %v\n", downloadCount, time.Since(startTime))
}

func TestUnzipAndDelete(t *testing.T) {
	path := filepath.Join(app.ModCacheDir(), testPkg1)
	zipPath := path + downloader.CUSTOM_ZIP_EXT

	err := fileutil.Unzip(zipPath, path, true)
	if err != nil {
		t.Fatalf("\nerror unpacking zip:\n\n%v", err)
	}

	t.Logf("\n\nSuccessfully unpacked and deleted:\n  %s", zipPath)
}
