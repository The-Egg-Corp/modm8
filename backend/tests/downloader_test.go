package backend

import (
	"fmt"
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"modm8/backend/thunderstore"
	"strings"
	"testing"
	"time"
)

const ZIP_EXT = ".zip"
const RAR_EXT = ".rar"

var TestDir = thunderstore.ModCacheDir("Lethal Company")

var testPool = downloader.DownloadPool{
	"https://thunderstore.io/package/download/Owen3H/CSync/3.0.1": fileutil.FileMetadata{
		Name:      "CSync-v3.0.1",
		Extension: ZIP_EXT,
	},
	"https://thunderstore.io/package/download/Owen3H/IntroTweaks/1.5.0": fileutil.FileMetadata{
		Name:      "Owen3H-IntroTweaks-1.5.0",
		Extension: RAR_EXT,
	},
	"https://thunderstore.io/package/download/Evaisa/LethalLib/0.16.1": fileutil.FileMetadata{
		Name:      "LethalLib-0.16.1",
		Extension: "",
	},
	"https://thunderstore.io/package/download/sfDesat/Orion/2.1.4": fileutil.FileMetadata{
		Name:      "ORION",
		Extension: ZIP_EXT,
	},
}

func TestDownloadMultipleFiles(t *testing.T) {
	errs, err := downloader.DownloadMultipleFiles(TestDir, testPool)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("\nFinished downloading files\n")

	if len(errs) > 0 {
		t.Log("\nEncountered non-fatal errors:")
		for _, e := range errs {
			fmt.Printf("		  %v\n", e)
		}
	}
}

// TODO: Implement in real backend when needed.
func BuildPackageDownloadURL(pkgName string) string {
	return tsDownloadDomain + strings.ReplaceAll(pkgName, "-", "/")
}

func TestDownloadFile(t *testing.T) {
	// Build URL and ticker before starting timer so it's *slightly* more accurate.
	url := BuildPackageDownloadURL(testPkg1)
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	startTime := time.Now()

	resp, err := downloader.DownloadFile(url, TestDir, fileutil.NewFileInfo(testPkg1, ZIP_EXT, TestDir))
	if err != nil {
		t.Skip("\n", err)
	}

Loop:
	for {
		select {
		case <-resp.Done:
			break Loop
		case <-ticker.C:
			t.Logf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(), resp.Size(), 100*resp.Progress(),
			)
		}
	}

	// Finished, check for errors
	if err := resp.Err(); err != nil {
		t.Fatalf("Download failed: %v\n", err)
	}

	elapsed := time.Since(startTime)
	t.Logf("\nDownload completed!\n\nSaved to: %v \nTook: %v\n", resp.Filename, elapsed)
}
