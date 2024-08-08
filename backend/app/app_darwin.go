//go:build darwin

package app

func setPriority() {
	return
}

func initCommands() {
	openCmd = &Command{name: "open"}
}
