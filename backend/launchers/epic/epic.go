package epic

import (
	"fmt"

	gocmd "github.com/go-cmd/cmd"
)

const launchCmd = "com.epicgames.launcher://apps/%s?action=launch&silent=true"

type EpicLauncher struct {
}

func NewRunner() *EpicLauncher {
	return &EpicLauncher{}
}

func (el *EpicLauncher) LaunchGame(id string) (*gocmd.Cmd, error) {
	return LaunchGame(id)
}

func LaunchGame(id string) (*gocmd.Cmd, error) {
	cmd := GetPlatformCommand(fmt.Sprintf(launchCmd, id))

	// TODO: Check if EGS supports launch args and implement them here.

	statusChan := cmd.Start()
	status := <-statusChan

	// Clean that bitch up properly.
	// See -> https://github.com/go-cmd/cmd?tab=readme-ov-file#proper-process-termination
	err := cmd.Stop()

	if status.Error != nil {
		err = status.Error
	}

	return cmd, err
}
