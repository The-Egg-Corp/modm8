package core

import (
	"context"
	"runtime"
)

type App struct {
	Settings *AppSettings `json:"settings"`
	Ctx      context.Context
}

func (a *App) GetSettings() *AppSettings {
	return a.Settings
}

func NewApp() *App {
	return &App{
		Settings: NewSettings(),
	}
}

// startup is called when the app starts.
// The context is saved so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
	a.Settings.Load()
}

func (a *App) Restart() {

}

func NumCPU() uint8 {
	return uint8(runtime.NumCPU())
}

func (a *App) NumCPU() uint8 {
	return NumCPU()
}

// func (a *App) Shutdown(ctx context.Context) {

// }

// Fetches the tag of the latest modm8 GitHub release.
//
// https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
