//go:build darwin

package app

func setPriority() error {
	return nil
}

func initOpenCommand() {
	openCmd = &Command{name: "open"}
}

func IsWindowsAdmin() bool {
	return false
}
