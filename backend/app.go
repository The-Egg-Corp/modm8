package backend

import (
	"context"
)

// Define App type
type App struct {
	Ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// startup is called when the app starts.
// The context is saved so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {

}

func CheckForUpdate() {

}
