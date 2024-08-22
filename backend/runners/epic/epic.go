package epic

import (
	"fmt"

	gocmd "github.com/go-cmd/cmd"
)

type EpicRunner struct {
}

func NewRunner() *EpicRunner {

	return &EpicRunner{}
}

func (runner *EpicRunner) LaunchGame(id string) (*gocmd.Cmd, error) {
	return LaunchGame(id)
}

func LaunchGame(id string) (*gocmd.Cmd, error) {
	uri := fmt.Sprintf("com.epicgames.launcher://apps/%s?action=launch&silent=true", id)
	cmd := GetPlatformCommand(uri)

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
