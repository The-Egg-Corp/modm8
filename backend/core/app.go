package core

import (
	"context"
	"errors"
	"os"
	"path"
	"runtime"

	"golang.org/x/sys/windows"

	gocmd "github.com/go-cmd/cmd"
	wRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Command struct {
	name string
	args *[]string
}

var (
	ErrUnsupportedPlatform = errors.New("unsupported platform")
	openCmd                *Command
)

type App struct {
	Ctx         context.Context
	Settings    *AppSettings `json:"settings"`
	Persistence *Persistence `json:"persistence"`
}

func (a *App) GetSettings() *AppSettings {
	return a.Settings
}

func (a *App) GetPersistence() *Persistence {
	return a.Persistence
}

func NewApp() *App {
	return &App{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
	}
}

// Called before anything else in main and before Wails runs.
//
// Any initial app setup logic should be done here.
func (a *App) Init() {
	a.Settings.Load()
	a.Persistence.Load()

	switch runtime.GOOS {
	case "windows":
		setWindowsPriority(windows.ABOVE_NORMAL_PRIORITY_CLASS)

		// Essentially replicates Electron's `shell.openExternal`
		openArgs := []string{"url.dll,FileProtocolHandler"}
		openCmd = &Command{
			name: "rundll32",
			args: &openArgs,
		}
	case "darwin":
		openCmd = &Command{name: "open"}
	case "linux":
		openCmd = &Command{name: "xdg-open"}
	}
}

// Called when the app starts. The context is saved so we can call the runtime methods.
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx

	if a.Persistence.Window.Maximized {
		wRuntime.WindowMaximise(ctx)
		return
	}

	//wRuntime.WindowSetPosition(ctx, a.Persistence.Window.X, a.Persistence.Window.Y)
}

// Called when the application is about to quit, either by clicking the window close button or calling runtime.Quit.
func (a *App) OnBeforeClose(ctx context.Context) bool {
	a.Persistence.ApplyCurrentWindowState(ctx)
	return false
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *App) Shutdown(ctx context.Context) {
	a.Persistence.Save()
}

func (a *App) Restart() {

}

func NumCPU() uint8 {
	return uint8(runtime.NumCPU())
}

func (a *App) NumCPU() uint8 {
	return NumCPU()
}

func ConfigDir() string {
	dir, _ := os.UserConfigDir()
	return path.Join(dir, "modm8")
}

func SettingsPath() string {
	return path.Join(ConfigDir(), "settings.toml")
}

func PersistencePath() string {
	return path.Join(ConfigDir(), "persistence.toml")
}

func (a *App) OpenExternal(path string) error {
	if openCmd == nil {
		return ErrUnsupportedPlatform
	}

	// Run cmd with the specified path as last argument
	cmd := gocmd.NewCmd(openCmd.name, append(*openCmd.args, path)...)
	statusChan := cmd.Start()

	status := <-statusChan
	if status.Error != nil {
		return status.Error
	}

	// Clean that bitch up (properly)
	// https://github.com/go-cmd/cmd?tab=readme-ov-file#proper-process-termination
	return cmd.Stop()
}

func setWindowsPriority(prio uint32) error {
	handle := windows.CurrentProcess()
	return windows.SetPriorityClass(handle, prio)
}

// Fetches the tag of the latest modm8 GitHub release.
//
// TODO: Poll this and return version string of latest -> https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
