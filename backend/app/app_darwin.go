//go:build darwin

package app

func setPriority() error {
	return nil
}

func initCommands() {
	openCmd = &Command{name: "open"}
}
