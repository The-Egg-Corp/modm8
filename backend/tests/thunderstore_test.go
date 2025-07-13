package backend

import (
	"modm8/backend/installing"
	"modm8/backend/loaders"
	"modm8/backend/thunderstore"
	"os"
	"testing"
	"time"

	v1 "github.com/the-egg-corp/thundergo/v1"
)

const tsDownloadDomain = "https://thunderstore.io/package/download/"
const testPkg1 = "sfDesat-Orion-2.1.4"

var comm = v1.Community{
	Identifier: "lethal-company",
}

var schema = thunderstore.NewThunderstoreSchema()

func TestGetEcosystem(t *testing.T) {
	ecosys, err := schema.GetEcosystem()
	if err != nil {
		t.Fatalf("GetEcosystem returned error:\n%v", err)
	}
	if ecosys == nil {
		t.Fatal("GetEcosystem returned nil ecosystem")
	}

	path, err := thunderstore.GetFallbackEcosystemPath()
	if err != nil {
		t.Fatalf("failed to get fallback path:\n%v", err)
	}

	if _, err := os.Stat(*path); err != nil {
		t.Fatalf("expected fallback file to exist at %s, but got error: %v", *path, err)
	}
}

func TestInstallWithDependencies(t *testing.T) {
	ecosys, err := schema.GetEcosystem()
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
