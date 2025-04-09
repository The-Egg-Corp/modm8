//go:build linux

package app

func setPriority() error {
	return nil
}

func initOpenCommand() {
	openCmd = &Command{name: "xdg-open"}
}
