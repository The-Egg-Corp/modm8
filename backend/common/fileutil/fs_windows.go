//go:build windows

package fileutil

import (
	"os/exec"
	"syscall"
)

// Silently runs link command with junction flag.
func SymlinkOrJunction(target, source string) error {
	cmd := exec.Command("cmd", "/C", "mklink", "/J", target, source)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}
