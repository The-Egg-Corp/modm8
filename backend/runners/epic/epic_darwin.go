//go:build darwin

package epic

import gocmd "github.com/go-cmd/cmd"

func GetPlatformCommand(uri string) *gocmd.Cmd {
	return gocmd.NewCmd("open", uri)
}
