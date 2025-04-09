//go:build windows

package app

import "golang.org/x/sys/windows"

func setPriority() error {
	handle := windows.CurrentProcess()
	return windows.SetPriorityClass(handle, windows.ABOVE_NORMAL_PRIORITY_CLASS)
}

// Essentially replicates Electron's `shell.openExternal`.
func initOpenCommand() {
	openArgs := []string{"url.dll,FileProtocolHandler"}
	openCmd = &Command{
		name: "rundll32",
		args: &openArgs,
	}
}
