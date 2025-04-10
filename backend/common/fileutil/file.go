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

func NewFileInfo(name string, extension string, parentDir string) FileMetadata {
	return FileMetadata{
		Name:      name,
		Extension: extension,
		ParentDir: parentDir,
	}
}

func (fi FileMetadata) GetPath() string {
	return filepath.Join(fi.ParentDir, fi.GetCombined())
}

// This method will attempt to combine and return this file's name and extension.
// If this file has no extension or is an empty/blank string, only the file name is returned.
// The correct format is ensured (fileName.ext) regardless of whether the extension starts with a period or not.
func (fi FileMetadata) GetCombined() string {
	if strings.TrimSpace(fi.Extension) == "" {
		return fi.Name
	}

	// Ensures correct format regardless of if the period was specified or not.
	ext, _ := strings.CutPrefix(fi.Extension, ".")
	return fi.Name + "." + ext
}
