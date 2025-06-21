//go:build unix

package fileutil

import (
	"os"
)

func SymlinkOrJunction(target, source string) error {
	return os.Symlink(source, target)
}
