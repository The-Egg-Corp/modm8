package backend

import (
	"fmt"
	"modm8/backend"
	"modm8/backend/common/profile"
	"os"
	"path/filepath"
	"testing"
)

func TestGetDirsAtPath(t *testing.T) {
	dirs, err := backend.GetDirsAtPath(profile.GameProfilesPath("Lethal Company"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dirs)
}

func TestLinkDir(t *testing.T) {
	TEMP := os.TempDir()
	source := filepath.Join(TEMP, "linkmod_source")
	target := filepath.Join(TEMP, "linkmod_target")

	os.RemoveAll(source)
	os.RemoveAll(target)

	if err := os.Mkdir(source, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}

	if _, err := os.Stat(target); err == nil {
		os.RemoveAll(target)
	}

	testFile := filepath.Join(source, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello"), 0644); err != nil {
		t.Fatalf("Failed to create test file in source: %v", err)
	}

	err := backend.LinkDir(target, source)
	if err != nil {
		t.Fatalf("LinkDir failed: %v", err)
	}

	info, err := os.Stat(target)
	if err != nil {
		t.Fatalf("Target does not exist after linking: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("Target exists but is not a directory")
	}

	// Verify link works by checking contents exist.
	linkedFile := filepath.Join(target, "test.txt")
	if _, err := os.Stat(linkedFile); err != nil {
		t.Fatalf("test.txt does not exist in target link: %v", err)
	}

	os.RemoveAll(source)
	os.RemoveAll(target)
}
