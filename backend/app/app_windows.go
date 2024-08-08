//go:build windows

package app

import "golang.org/x/sys/windows"

func setPriority() error {
	handle := windows.CurrentProcess()
	return windows.SetPriorityClass(handle, windows.ABOVE_NORMAL_PRIORITY_CLASS)
}

func initCommands() {
	// Essentially replicates Electron's `shell.openExternal`
	openArgs := []string{"url.dll,FileProtocolHandler"}
	openCmd = &Command{
		name: "rundll32",
		args: &openArgs,
	}
}
