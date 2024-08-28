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
		if d.IsDir() {
			return nil
		}

		for _, s := range exts {
			if strings.HasSuffix(path, "."+s) {
				files = append(files, path)
				return nil
			}
		}

		return nil
	})
	return files, err
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
