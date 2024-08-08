//go:build linux

package app

func setPriority() {
	return
}

func initCommands() {
	openCmd = &Command{name: "xdg-open"}
}
