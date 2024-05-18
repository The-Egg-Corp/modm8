package core

import (
	"context"
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

// func (a *App) Shutdown(ctx context.Context) {

// }

func CheckForUpdate() {

}
