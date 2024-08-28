package downloader

import (
	"fmt"
	"modm8/backend"
	"modm8/backend/common/fileutil"
	"path/filepath"

	"github.com/cavaliergopher/grab/v3"
)

type DownloadRequest struct {
	URL        string
	OutputFile fileutil.FileInfo
}

func NewDownloadRequest(url string, fileName string, extension *string) DownloadRequest {
	return DownloadRequest{
		URL:        url,
		OutputFile: fileutil.NewFileInfo(fileName, extension),
	}
}

func DownloadZip(url, dirPath, fileName string) (*grab.Response, error) {
	ext := ".zip"
	return DownloadFile(url, dirPath, fileutil.NewFileInfo(fileName, &ext))
}

// Makes a GET request to a download URL and saves it to the specified directory (created if it doesn't exist).
//
// Notes:
//
// - The directory path is normalized and cleaned to be platform-independent.
//
// - While the download is in progress, the file is renamed with the specified extension (if not nil) plus ".tmp".
func DownloadFile(url, dirPath string, fi fileutil.FileInfo) (*grab.Response, error) {
	outputPath := filepath.Join(filepath.Clean(dirPath), fi.GetCombined())
	if exists, _ := backend.ExistsAtPath(outputPath); exists {
		return nil, fmt.Errorf("file already exists: %s", outputPath)
	}

	res, err := grab.Get(outputPath, url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	return res, nil
}

// Calls `DownloadFile` concurrently for the given urls, putting them all into the same specified path.
func DownloadMultipleFiles(destPath string, requests ...DownloadRequest) {

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
