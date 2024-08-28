package backend

import (
	"modm8/backend/common/downloader"
	"modm8/backend/common/fileutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const testPath = "modm8\\Games\\LethalCompany\\Profiles\\test"
const testPkg1 = "Owen3H-CSync-3.0.1"

//const testPkg2 = "Owen3H-IntroTweaks-1.5.0"

//var testPkgs = []string{testPkg1, testPkg2}

func GetTestDir() string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, testPath)
}

func TestInstallPackage(t *testing.T) {

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

// func TestUnpackZip(t *testing.T) {
// 	path := filepath.Join(GetTestDir(), testPkg1)
// 	zipPath := path + ".zip"

// 	// Initialize an extractor
// 	e, err := fastzip.NewExtractor(zipPath, path)
// 	if err != nil {
// 		t.Fatalf("\nerror creating extractor:\n\n%v", err)
// 	}

// 	// Use extractor to extract the archive. (As a background operation)
// 	if err = e.Extract(context.Background()); err != nil {
// 		t.Fatalf("\nerror extracting zip:\n\n%v", err)
// 		e.Close()
// 	}

// 	// At this point, extraction is complete.
// 	t.Logf("\nSuccessfully extracted %s", zipPath)
// 	e.Close()

// 	// Cleanup by deleting zip
// 	if err := os.Remove(zipPath); err != nil {
// 		t.Fatalf("\nerror deleting zip file:\n\n%v", err)
// 	}

// 	t.Logf("\nSuccessfully deleted %s", zipPath)
// }

func TestDownloadZip(t *testing.T) {
	url := "https://thunderstore.io/package/download/" + strings.ReplaceAll(testPkg1, "-", "/")
	ext := ".zip"

	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	startTime := time.Now()

	resp, err := downloader.DownloadFile(url, GetTestDir(), fileutil.NewFileInfo(testPkg1, &ext))

	if err != nil {
		t.Skip("\n", err)
	}

	var timeTaken time.Duration

Loop:
	for {
		select {
		case <-resp.Done:
			timeTaken = time.Since(startTime)
			break Loop
		case <-ticker.C:
			t.Logf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(), resp.Size(), 100*resp.Progress(),
			)
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		t.Fatalf("Download failed: %v\n", err)
	}

	t.Logf("\nDownload completed!\n\nSaved to: %v \nTook: %v\n", resp.Filename, timeTaken)
}

// func TestDownloadZip2(t *testing.T) {
// 	fmt.Println("Download Started")

// 	start := time.Now()
// 	err := DownloadFile(
// 		"D:\\Temp\\Downloads\\large.dat",
// 		"https://proof.ovh.net/files/100Mb.dat",
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("Download Finished. Took %v", time.Since(start))
// }

// // WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// // and we can pass this into io.TeeReader() which will report progress on each write cycle.
// type WriteCounter struct {
// 	Total uint64
// }

// func (wc *WriteCounter) Write(p []byte) (int, error) {
// 	n := len(p)
// 	wc.Total += uint64(n)

// 	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))

// 	return n, nil
// }

// // DownloadFile will download a url to a local file. It's efficient because it will
// // write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// // into Copy() to report progress on the download.
// func DownloadFile(filepath string, url string) error {
// 	// Create the file, but give it a tmp file extension, this means we won't overwrite a
// 	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
// 	out, err := os.Create(filepath + ".tmp")
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		out.Close()
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create our progress reporter and pass it to be used alongside our writer
// 	counter := &WriteCounter{}
// 	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
// 		out.Close()
// 		return err
// 	}

// 	fmt.Print("\n") // The progress use the same line so print a new line once it's finished downloading
// 	out.Close()     // Close the file without defer so it can happen before Rename()

// 	if err = os.Rename(filepath+".tmp", filepath); err != nil {
// 		return err
// 	}

// 	return nil
// }
