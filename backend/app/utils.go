package app

import (
	"modm8/backend/common/fileutil"
	"path/filepath"
)

// This struct is bound to Wails and is responsible for providing utility functions relating to
// the file system, un/zipping, Unity Doorstop and other general purpose functions.
type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) ExistsInDir(dir, item string) (bool, error) {
	return fileutil.ExistsInDir(dir, item)
}

func (u *Utils) ExistsAtPath(path string, clean bool) (bool, error) {
	if clean {
		path = filepath.Clean(path)
	}

	return fileutil.ExistsAtPath(path)
}

func (u *Utils) GetDirsAtPath(path string, exts []string) ([]string, error) {
	return fileutil.GetDirsAtPath(path)
}

func (u *Utils) GetFilesWithExts(path string, exts []string) ([]string, error) {
	return fileutil.GetFilesWithExts(path, exts)
}

func (u *Utils) GetFilesInZip(data []byte) (map[string][]byte, error) {
	return fileutil.GetFilesInZip(data)
}
