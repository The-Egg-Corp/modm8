package backend

import (
	"fmt"
	"modm8/backend/common/fileutil"
	"modm8/backend/common/profile"
	"os"
	"path/filepath"
	"testing"
)

func TestGetDirsAtPath(t *testing.T) {
	dirs, err := fileutil.GetDirsAtPath(profile.GameProfilesPath("Lethal Company"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dirs)
}

func TestLinkDirSucceeds(t *testing.T) {
	TEMP := os.TempDir()
	source := filepath.Join(TEMP, "linkmod_source")
	target := filepath.Join(TEMP, "linkmod_target")

	os.RemoveAll(source)
	os.RemoveAll(target)

	//#region Pre-checks
	if err := os.Mkdir(source, os.ModePerm); err != nil {
		t.Fatalf("failed to create source dir:\n%v", err)
	}

	testFile := filepath.Join(source, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello"), os.ModePerm); err != nil {
		t.Fatalf("failed to create test file in source:\n%v", err)
	}
	//#endregion

	// Actual test
	err := fileutil.LinkDir(target, source)
	if err != nil {
		t.Fatalf("failed to link dir:\n%v", err)
	}

	//#region Post-checks
	info, err := os.Stat(target)
	if err != nil {
		t.Fatalf("target does not exist after linking:\n%v", err)
	}
	if !info.IsDir() {
		t.Fatalf("target exists but is not a directory")
	}

	linkedFile := filepath.Join(target, "test.txt")
	if _, err := os.Stat(linkedFile); err != nil {
		t.Fatalf("test.txt does not exist in target link:\n%v", err)
	}
	//#endregion

	os.RemoveAll(source)
	os.RemoveAll(target)
}

func TestLinkDirFails(t *testing.T) {
	TEMP := os.TempDir()
	source := filepath.Join(TEMP, "linkmod_source")
	target := filepath.Join(TEMP, "linkmod_target")

	os.RemoveAll(source)
	os.RemoveAll(target)

	// Should fail with source directory error.
	err := fileutil.LinkDir(target, source)
	if err != nil {
		t.Fatalf("failed to link dir:\n%v", err)
	}
}
