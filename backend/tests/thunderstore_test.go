package backend

import (
	"modm8/backend/common/fileutil"
	"modm8/backend/thunderstore"
	"path/filepath"
	"strings"
	"testing"
	"time"

	v1 "github.com/the-egg-corp/thundergo/v1"
)

const tsDownloadDomain = "https://thunderstore.io/package/download/"
const testPkg1 = "sfDesat-Orion-2.1.4"

func GetTestDir() string {
	//cacheDir, _ := os.UserConfigDir()
	return filepath.Join("H:\\", "Program Files", "modm8", "Thunderstore", "LethalCompany", "ModCache")
}

func TestPathJoin(t *testing.T) {
	str := GetTestDir()
	t.Log(str)

	if !strings.HasPrefix(str, "H") {
		t.Fatal("dir path is relative. should be absolute")
	}
}

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

	thunderstore.CurModCacheDir = thunderstore.ModCacheDir("LethalCompany")
	thunderstore.InstallWithDependencies(pkg.LatestVersion(), pkgs, &errs, &downloadCount)

	t.Logf("\nDownloaded %v packages in %v\n", downloadCount, time.Since(startTime))
}

func TestUnzip(t *testing.T) {
	path := filepath.Join(GetTestDir(), testPkg1)
	zipPath := path + ".zip"

	err := fileutil.Unzip(zipPath, path, true)
	if err != nil {
		t.Fatalf("\nerror unpacking zip:\n\n%v", err)
	}

	t.Logf("\n\nSuccessfully unpacked and deleted:\n  %s", zipPath)
}
