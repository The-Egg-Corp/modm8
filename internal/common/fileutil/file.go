package fileutil

import (
	"path/filepath"
	"strings"
)

type FileMetadata struct {
	Name      string
	Extension string
	ParentDir string
}

func NewFileInfo(parentDir string, name string, extension string) FileMetadata {
	return FileMetadata{
		ParentDir: parentDir,
		Name:      name,
		Extension: extension,
	}
}

func (fi FileMetadata) Path() string {
	return filepath.Join(fi.ParentDir, fi.NameAndExt())
}

// This method will attempt to combine and return this file's name and extension.
// If this file has no extension or is an empty/blank string, only the file name is returned.
// The correct format is ensured (fileName.ext) regardless of whether the extension starts with a period or not.
func (fi FileMetadata) NameAndExt() string {
	if strings.TrimSpace(fi.Extension) == "" {
		return fi.Name
	}

	// Ensures correct format regardless of if the period was specified or not.
	ext, _ := strings.CutPrefix(fi.Extension, ".")
	return fi.Name + "." + ext
}
