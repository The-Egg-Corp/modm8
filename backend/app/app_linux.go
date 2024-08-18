//go:build linux

package app

func setPriority() error {
	return nil
}

func initCommands() {
	openCmd = &Command{name: "xdg-open"}
}
