package fileutil

import (
	"path/filepath"
	"strings"
)

type FileMetadata struct {
	Name      string
	Extension *string
	ParentDir string
}

func NewFileInfo(name string, extension *string, parentDir string) FileMetadata {
	return FileMetadata{
		Name:      name,
		Extension: extension,
		ParentDir: parentDir,
	}
}

func (fi FileMetadata) GetPath() string {
	return filepath.Join(fi.ParentDir, fi.GetCombined())
}

func (fi FileMetadata) GetCombined() string {
	if fi.Extension == nil {
		return fi.Name
	}

	// Ensures correct format regardless of if the period was specified or not.
	ext, _ := strings.CutPrefix(*fi.Extension, ".")
	return fi.Name + "." + ext
}
