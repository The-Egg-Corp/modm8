package fileutil

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
func ExistsInDir(targetDir string, item string) (bool, error) {
	path := filepath.Join(filepath.Clean(targetDir), item)
	return ExistsAtPath(path)
}

// Checks whether something exists at the given absolute path (dir or file).
// This func will follow symbolic links and return true if the link source is existent.
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
	// Symlink does this, but we can exit even earlier and provide better error context.
	sfi, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("error validating source path. dir may not exist:\n%v", err)
	}

	if !sfi.IsDir() {
		return fmt.Errorf("invalid source path. must be a dir")
	}

	return CreateSymlinkOrJunction(target, source)
}

// Scans the contents of the file at the given path line by line. At the first instance of a non-blank line,
// the line is returned with no line endings or carriage return, making this function platform-independent.
func FindFirstValidLine(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lineContent := strings.TrimSpace(s.Text())
		if lineContent != "" {
			return &lineContent, nil // Scan should have already removed \n and \r.
		}
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error reading file at %s\n%v", path, err)
	}

	return nil, fmt.Errorf("no valid line found in file: %s", path)
}

// Like [FindFirstValidLine], this function scans the contents of the file at the given path line by line,
// but also adds an extra constraint where the line must be in a proper semver format.
//
// In summary, the first line that is both non-blank and valid semver will be returned.
func FindFirstSemverLine(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lineContent := strings.TrimSpace(s.Text())
		if lineContent == "" {
			continue
		}

		semverParts := strings.Split(lineContent, ".")
		if len(semverParts) != 3 {
			continue
		}

		if _, err1 := strconv.Atoi(semverParts[0]); err1 != nil {
			continue
		}
		if _, err2 := strconv.Atoi(semverParts[1]); err2 != nil {
			continue
		}
		if _, err3 := strconv.Atoi(semverParts[2]); err3 != nil {
			continue
		}

		return &lineContent, nil
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error reading file at %s\n%v", path, err)
	}

	return nil, fmt.Errorf("no semver line found in file: %s", path)
}
