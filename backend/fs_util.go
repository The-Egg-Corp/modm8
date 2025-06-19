package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// Takes a list of paths and returns their base file/folder names. Example:
//
// In: ["C:/Games/Game", "C:/SomeFolder/myfile.json"]
//
// Out: ["Game", "myfile.json"]
func GetBaseNames(paths []string) []string {
	names := make([]string, len(paths))
	for i, path := range paths {
		names[i] = filepath.Base(path)
	}

	return names
}

func GetDirsAtPath(path string) ([]string, error) {
	var dirs []string
	var name = filepath.Base(path)

	err := filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
		if entry.IsDir() && entry.Name() != name {
			dirs = append(dirs, path)
		}

		return nil
	})

	return dirs, err
}

func GetFilesWithExts(path string, exts []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
		if !entry.IsDir() {
			for _, s := range exts {
				if strings.HasSuffix(path, "."+s) {
					files = append(files, path)
					break
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

// The same as os.ReadFile, but the path is cleaned automatically.
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(path))
}

// The same as os.WriteFile, but the path is cleaned automatically.
func WriteFile(path string, data []byte) error {
	return os.WriteFile(filepath.Clean(path), data, os.ModePerm)
}

// The same as os.Mkdir, but the path is cleaned automatically and perm is os.ModePerm.
func MkDir(path string) error {
	return os.Mkdir(filepath.Clean(path), os.ModePerm)
}

// The same as os.MkdirAll, but the path is cleaned automatically and perm is os.ModePerm.
func MkDirAll(path string) error {
	return os.MkdirAll(filepath.Clean(path), os.ModePerm)
}

// Links directory `target` to directory `source` using a Symlink (Junction on Windows), essentially making `target` mirror its contents.
//
// NOTE: While we can use [os.Symlink] on Windows, we don't currently as it would require admin privileges which becomes a massive pain in the ass.
func LinkDir(target, source string) error {
	// Symlink will do this, but this will exit even earlier and is more informative than cmd.Run()
	sfi, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("error validating source path. dir may not exist:\n%v", err)
	}

	if !sfi.IsDir() {
		return fmt.Errorf("invalid source path. must be a dir")
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "mklink", "/J", target, source)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		return cmd.Run()
	}

	return os.Symlink(source, target)
}
