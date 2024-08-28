package fileutil

import (
	"context"
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
