package app

import (
	"context"
	"errors"
	"modm8/backend/app/appcore"
	"modm8/backend/app/appservices"

	gocmd "github.com/go-cmd/cmd"
	wuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	ErrUnsupportedPlatform = errors.New("unsupported platform")
	openCmd                *Command
)

type Command struct {
	name string
	args *[]string
}

type Application struct {
	WailsCtx context.Context
	core     *appcore.AppCore // TODO: Maybe just rename to Env and make Utils a service.
	services *appservices.AppServices
}

func NewApplication() *Application {
	core := appcore.NewAppCore()
	services := appservices.NewAppServices(core)

	return &Application{
		WailsCtx: context.TODO(),
		core:     core,
		services: services,
	}
}

func (app *Application) GetWailsContext() context.Context {
	return app.WailsCtx
}

func (app *Application) GetCore() *appcore.AppCore {
	return app.core
}

func (app *Application) GetServices() *appservices.AppServices {
	return app.services
}

func (app *Application) GetSettings() *appcore.AppSettings {
	return app.GetCore().Settings
}

func (app *Application) GetPersistence() *appcore.Persistence {
	return app.GetCore().Persistence
}

func (app *Application) GetUtils() *appcore.Utils {
	return app.GetCore().Utils
}

func (app *Application) IsWindowsAdmin() bool {
	return IsWindowsAdmin()
}

// Called before anything else in main and before Wails runs.
//
// Any initial app setup logic should be done here.
func (app *Application) Init() (errs []error) {
	err := app.GetSettings().Load()
	if err != nil {
		errs = append(errs, err)
	}

	err = app.GetPersistence().Load()
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
func (app *Application) Startup(wailsCtx context.Context) {
	app.WailsCtx = wailsCtx

	if app.GetPersistence().WindowState.Maximized {
		wuntime.WindowMaximise(wailsCtx)
		return
	}

	//wuntime.WindowSetPosition(ctx, a.Persistence.Window.X, a.Persistence.Window.Y)
}

// Called when the application is about to quit, either by clicking the window close button or calling runtime.Quit.
func (app *Application) OnBeforeClose(ctx context.Context) bool {
	app.GetPersistence().ApplyCurrentWindowState(ctx)
	return false
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *Application) Shutdown(ctx context.Context) {
	a.GetPersistence().Save()
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

// Fetches the tag of the latest modm8 GitHub release.
//
// TODO: Poll this and return version string of latest -> https://api.github.com/repos/The-Egg-Corp/modm8/releases/latest
func LatestReleaseVersion() {

}
