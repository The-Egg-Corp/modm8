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

		defer rc.Close()

		// Read the file contents
		contents, err := io.ReadAll(rc)
		if err != nil {
			return nil, fmt.Errorf("failed to read file contents: %w", err)
		}

		// Store the file in the map
		files[f.Name] = contents
	}

	return files, nil
}
