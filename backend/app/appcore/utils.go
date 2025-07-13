package appcore

import (
	"modm8/backend/common/fileutil"
	"path/filepath"
	"runtime"

	"github.com/samber/lo"
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

func (u *Utils) NumCPU() uint8 {
	return NumCPU()
}

func (u *Utils) GetMaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

func NumCPU() uint8 {
	return uint8(runtime.NumCPU())
}

// Sets GOMAXPROCS to given value and ensures it is clamped between 1 and NumCPU*2 as any further may degrade performance due to context switching.
// Note that blocking syscalls can have their own threads regardless of the limit set here.
func SetMaxProcs(num uint8) int {
	return runtime.GOMAXPROCS(lo.Clamp(int(num), 1, runtime.NumCPU()*2))
}
