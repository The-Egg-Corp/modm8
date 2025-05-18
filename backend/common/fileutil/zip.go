package fileutil

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/saracen/fastzip"
)

// Opens a zip file, extracts its contents via a background operation,
// then subsequently deletes the no longer useful zip.
func Unzip(path, dest string, delete bool) error {
	// Initialize an extractor
	e, err := fastzip.NewExtractor(path, dest)
	if err != nil {
		return err
	}
	defer e.Close()

	// Use extractor to extract the archive. (As a background operation)
	if err = e.Extract(context.Background()); err != nil {
		return err
	}

	// Close before deleting to avoid contention.
	err = e.Close()
	if err == nil && delete {
		// We no longer need the original zip.
		if err = os.Remove(path); err != nil {
			return err
		}
	}

	return err
}

// Given a byte slice, which assumed to be data representing a zip file, this function will attempt to
// read said zip, scanning every file and collecting its contents into a map where the key represents
// the file name and the value represents the file contents as a byte slice.
//
// Note that a byte slice (file) in the map may be nil if an error occurred or the file is empty.
func GetFilesInZip(data []byte) (map[string][]byte, error) {
	// Create a zip reader
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("failed to create zip reader: %w", err)
	}

	files := make(map[string][]byte)
	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file in zip: %w", err)
		}

		// Read the file contents
		contents, err := io.ReadAll(rc)
		rc.Close()

		if err != nil {
			return nil, fmt.Errorf("failed to read file contents: %w", err)
		}

		// Store the file in the map
		files[f.Name] = contents
	}

	return files, nil
}
