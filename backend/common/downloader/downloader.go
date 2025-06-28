package downloader

import (
	"errors"
	"fmt"
	"modm8/backend/common/fileutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/cavaliergopher/grab/v3"
	"golang.org/x/sync/errgroup"
)

const CUSTOM_ZIP_EXT = ".m8z"

// Alias for a map, where the key is the download URL and value is the associated output file info.
type DownloadPool = map[string]fileutil.FileMetadata

func DownloadZip(url, dirPath, fileName string) (*grab.Response, error) {
	return DownloadFile(url, dirPath, fileutil.NewFileInfo(fileName, CUSTOM_ZIP_EXT, dirPath))
}

// Makes a GET request to a download URL and saves it to the specified directory (created if it doesn't exist).
//
// The directory path is normalized and cleaned to be platform-independent.
func DownloadFile(url, dirPath string, fi fileutil.FileMetadata) (*grab.Response, error) {
	outputPath := filepath.Join(filepath.Clean(dirPath), fi.GetCombined())
	if exists, _ := fileutil.ExistsAtPath(outputPath); exists {
		return nil, fmt.Errorf("file/dir already exists: %s", outputPath)
	}

	res, err := grab.Get(outputPath, url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	return res, nil
}

// Downloads all files in the pool concurrently, putting them all into the same specified path.
//
// The queue map allows us to customize the output file info per URL.
func DownloadMultipleFiles(destPath string, pool DownloadPool) (map[string]error, error) {
	fi, err := os.Stat(destPath)
	if err != nil {
		return nil, err
	}

	if !fi.IsDir() {
		return nil, errors.New("destination is not a directory")
	}

	if len(pool) < 1 {
		return nil, errors.New("error downloading: specified queue is empty")
	}

	var g errgroup.Group

	// Slice to collect errors
	errs := make(map[string]error)
	errsMut := &sync.Mutex{}

	for url, outputFile := range pool {
		// Capture the vars instead of wrapping in a func
		url, outputFile := url, outputFile
		g.Go(func() error {
			// Call the DownloadFile function for each URL
			_, err := DownloadFile(url, destPath, outputFile)
			if err != nil {
				// Lock the map before writing to avoid race conditions
				errsMut.Lock()
				errs[url] = err
				errsMut.Unlock()
			}

			return nil
		})
	}

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return errs, err // Return the map with errors and the first error encountered by the group
	}

	return errs, nil
}

// ================================ ARCHIVED =====================================
// // ProgressChannel is a channel for reporting progress updates as integers.
// type ProgressChannel chan int

// // ProgressWriter tracks the progress of downloading and saving.
// type ProgressWriter struct {
// 	Total      int64
// 	Written    int64
// 	ProgressCh ProgressChannel
// }

// // Write method to satisfy the io.Writer interface and report progress via the channel.
// func (pw *ProgressWriter) Write(p []byte) (int, error) {
// 	n := len(p)
// 	pw.Written += int64(n)

// 	if pw.Total > 0 {
// 		percentage := int(float64(pw.Written) / float64(pw.Total) * 100)
// 		if pw.ProgressCh != nil {
// 			pw.ProgressCh <- percentage
// 		}
// 	}

// 	return n, nil
// }
// ===============================================================================
