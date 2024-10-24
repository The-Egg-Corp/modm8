package backend

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func WalkDirExt(root string, exts []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			for _, s := range exts {
				if strings.HasSuffix(path, "."+s) {
					files = append(files, path)
					return nil
				}
			}
		}

		return nil
	})

	return files, err
}

// Platform-independent way of checking a file/dir exists in a directory.
func ExistsInDir(dir string, item string) (bool, error) {
	path := filepath.Join(filepath.Clean(dir), item)
	return ExistsAtPath(path)
}

func ExistsAtPath(absPath string) (bool, error) {
	_, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false, nil
	}

	return err == nil, err
}

func ReadFile(path string) (*string, error) {
	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, nil
	}

	out := string(content)
	return &out, nil
}

func ContainsEqualFold(arr []string, item string) bool {
	lowerItem := strings.ToLower(item)

	for _, str := range arr {
		if strings.ToLower(str) == lowerItem {
			return true
		}
	}

	return false
}
