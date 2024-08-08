package app

import (
	"context"
	"errors"
	"os"
	"path"
	"runtime"

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

type Application struct {
	Ctx         context.Context
	Settings    *AppSettings `json:"settings"`
	Persistence *Persistence `json:"persistence"`
}

func (app *Application) GetSettings() *AppSettings {
	return app.Settings
}

func (app *Application) GetPersistence() *Persistence {
	return app.Persistence
}

func NewApp() *Application {
	return &Application{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
	}
}

// Called before anything else in main and before Wails runs.
//
// Any initial app setup logic should be done here.
func (app Application) Init() (errs []error) {
	var err error

	err = app.Settings.Load()
	if err != nil {
		errs = append(errs, err)
	}

	err = app.Persistence.Load()
	if err != nil {
		errs = append(errs, err)
	}

	//#region Platform specific
	initCommands()

	err = setPriority()
	if err != nil {
		errs = append(errs, err)
	}
	//#endregion

	return errs
}

// Called when the app starts. The context is saved so we can call the runtime methods.
func (app *Application) Startup(ctx context.Context) {
	app.Ctx = ctx

	if app.Persistence.Window.Maximized {
		wRuntime.WindowMaximise(ctx)
		return
	}

	//wRuntime.WindowSetPosition(ctx, a.Persistence.Window.X, a.Persistence.Window.Y)
}

// Called when the application is about to quit, either by clicking the window close button or calling runtime.Quit.
func (app *Application) OnBeforeClose(ctx context.Context) bool {
	app.Persistence.ApplyCurrentWindowState(ctx)
	return false
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *Application) Shutdown(ctx context.Context) {
	a.Persistence.Save()
}

func (a *Application) Restart() {

}

func (app *Application) OpenExternal(path string) error {
	if openCmd == nil {
		return ErrUnsupportedPlatform
	}

	// Run cmd with the specified path as last argument
	cmd := gocmd.NewCmd(openCmd.name, append(*openCmd.args, path)...)

	statusChan := cmd.Start()
	status := <-statusChan

	// Clean that bitch up properly.
	// See -> https://github.com/go-cmd/cmd?tab=readme-ov-file#proper-process-termination
	err := cmd.Stop()

	if status.Error != nil {
		err = status.Error
	}

	return err
}

func (app *Application) NumCPU() uint8 {
	return NumCPU()
}

func NumCPU() uint8 {
	return uint8(runtime.NumCPU())
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

// Fetches the tag of the latest modm8 GitHub release.
//
// TODO: Poll this and return version string of latest -> https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
