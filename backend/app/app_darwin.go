//go:build darwin

package app

func initOpenCommand() {
	openCmd = &Command{name: "open"}
}

func setPriority() error {
	return nil
}

func IsWindowsAdmin() bool {
	return false
}
