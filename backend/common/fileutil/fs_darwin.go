//go:build darwin

package fileutil

import (
	"fmt"
	"os"
)

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

	return os.Symlink(source, target)
}
