package app

import (
	"modm8/backend"
	"modm8/backend/common/fileutil"
	"path/filepath"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) ExistsInDir(dir, item string) (bool, error) {
	return backend.ExistsInDir(dir, item)
}

func (u Utils) ExistsAtPath(path string, clean bool) (bool, error) {
	if clean {
		path = filepath.Clean(path)
	}

	return backend.ExistsAtPath(path)
}

func (u *Utils) GetDirsAtPath(path string, exts []string) ([]string, error) {
	return backend.GetDirsAtPath(path)
}

func (u *Utils) GetFilesWithExts(path string, exts []string) ([]string, error) {
	return backend.GetFilesWithExts(path, exts)
}

func (u *Utils) GetFilesInZip(data []byte) (map[string][]byte, error) {
	return fileutil.GetFilesInZip(data)
}
