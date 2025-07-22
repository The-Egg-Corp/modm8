//go:build linux

package app

func initOpenCommand() {
	openCmd = &Command{name: "xdg-open"}
}

func setPriority() error {
	return nil
}

func IsWindowsAdmin() bool {
	return false
}
