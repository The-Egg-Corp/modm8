//go:build windows

package epic

import gocmd "github.com/go-cmd/cmd"

func GetPlatformCommand(uri string) *gocmd.Cmd {
	return gocmd.NewCmd("cmd", "/C", "start", uri)
}
