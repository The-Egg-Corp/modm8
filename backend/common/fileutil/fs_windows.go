//go:build windows

package fileutil

import (
	"os/exec"
	"syscall"
)

// Silently runs link command with junction flag.
func CreateSymlinkOrJunction(target, source string) error {
	cmd := exec.Command("cmd", "/C", "mklink", "/J", target, source)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}
