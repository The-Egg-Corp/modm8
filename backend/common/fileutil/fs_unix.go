//go:build unix

package fileutil

import (
	"os"
)

func CreateSymlinkOrJunction(target, source string) error {
	return os.Symlink(source, target)
}
