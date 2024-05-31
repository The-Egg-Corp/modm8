package core

import (
	"context"
	"os"
	"path"
	"runtime"

	wRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Ctx         context.Context
	Settings    *AppSettings `json:"settings"`
	Persistence *Persistence `json:"persistence"`
}

func (a *App) GetSettings() *AppSettings {
	return a.Settings
}

func NewApp() *App {
	return &App{
		Settings:    NewSettings(),
		Persistence: NewPersistence(),
	}
}

// Called when the app starts. The context is saved so we can call the runtime methods.
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx

	wRuntime.WindowSetPosition(ctx, a.Persistence.Window.X, a.Persistence.Window.Y)
}

// Called when the application is about to quit, either by clicking the window close button or calling runtime.Quit.
func (a *App) OnBeforeClose(ctx context.Context) bool {
	a.Persistence.SaveCurrentWindowState(ctx)
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

// Fetches the tag of the latest modm8 GitHub release.
//
// https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
