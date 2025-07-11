package app

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"

	gocmd "github.com/go-cmd/cmd"
	"github.com/samber/lo"
	wuntime "github.com/wailsapp/wails/v2/pkg/runtime"
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
	Utils       *Utils       `json:"utils"`
	Settings    *AppSettings `json:"settings"`
	Persistence *Persistence `json:"persistence"`
}

func (app *Application) GetSettings() *AppSettings {
	return app.Settings
}

func (app *Application) GetPersistence() *Persistence {
	return app.Persistence
}

func (app *Application) IsWindowsAdmin() bool {
	return IsWindowsAdmin()
}

func NewApp() *Application {
	return &Application{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
		Utils:       NewUtils(),
	}
}

// Called before anything else in main and before Wails runs.
//
// Any initial app setup logic should be done here.
func (app Application) Init() (errs []error) {
	err := app.Settings.Load()
	if err != nil {
		errs = append(errs, err)
	}

	err = app.Persistence.Load()
	if err != nil {
		errs = append(errs, err)
	}

	//#region Platform specific
	initOpenCommand()

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
		wuntime.WindowMaximise(ctx)
		return
	}

	//wuntime.WindowSetPosition(ctx, a.Persistence.Window.X, a.Persistence.Window.Y)
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

// TODO: Implement this so app can be restarted (like me) by itself.
func (a *Application) Restart() {

}

func (app *Application) OpenExternal(path string) error {
	if openCmd == nil {
		return ErrUnsupportedPlatform
	}

	// Run cmd with the specified path as last argument
	cmd := gocmd.NewCmd(openCmd.name, append(*openCmd.args, path)...)
	status := <-cmd.Start()

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

func (app *Application) GetMaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

// Sets GOMAXPROCS to given value and ensures it is clamped between 1 and NumCPU*2 as any further may degrade performance due to context switching.
// Note that blocking syscalls can have their own threads regardless of the limit set here.
func SetMaxProcs(num uint8) int {
	return runtime.GOMAXPROCS(lo.Clamp(int(num), 1, runtime.NumCPU()*2))
}

func ConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, "modm8")
}

// Uses the users config dir and returns a path to the mod cache.
func ModCacheDir() string {
	cacheDir, _ := os.UserConfigDir()
	return filepath.Join(cacheDir, "modm8", "ModCache")
}

func KeyPath() string {
	return filepath.Join(ConfigDir(), "nex.key")
}

func SettingsPath() string {
	return filepath.Join(ConfigDir(), "settings.toml")
}

func PersistencePath() string {
	return filepath.Join(ConfigDir(), "persistence.toml")
}

// Fetches the tag of the latest modm8 GitHub release.
//
// TODO: Poll this and return version string of latest -> https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
