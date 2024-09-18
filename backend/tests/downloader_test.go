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

var zip = ".zip"
var rar = ".rar"

var TestDir = thunderstore.ModCacheDir("Lethal Company")

var testPool = downloader.DownloadPool{
	"https://thunderstore.io/package/download/Owen3H/CSync/3.0.1": fileutil.FileMetadata{
		Name:      "CSync-v3.0.1",
		Extension: &zip,
	},
	"https://thunderstore.io/package/download/Owen3H/IntroTweaks/1.5.0": fileutil.FileMetadata{
		Name:      "Owen3H-IntroTweaks-1.5.0",
		Extension: &rar,
	},
	"https://thunderstore.io/package/download/Evaisa/LethalLib/0.16.1": fileutil.FileMetadata{
		Name:      "LethalLib-0.16.1",
		Extension: nil,
	},
	"https://thunderstore.io/package/download/sfDesat/Orion/2.1.4": fileutil.FileMetadata{
		Name:      "ORION",
		Extension: &zip,
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

func TestDownloadFile(t *testing.T) {
	url := tsDownloadDomain + strings.ReplaceAll(testPkg1, "-", "/")
	ext := "zip"

	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	startTime := time.Now()

	resp, err := downloader.DownloadFile(url, TestDir, fileutil.NewFileInfo(testPkg1, &ext, TestDir))
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

	t.Logf("\nDownload completed!\n\nSaved to: %v \nTook: %v\n", resp.Filename, time.Since(startTime))
}
