package app

import (
	"context"
	"errors"
	"modm8/backend/app/appctx"
	"modm8/backend/game"
	"modm8/backend/launchers/steam"
	"modm8/backend/profile"
	"modm8/backend/thunderstore"

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
	Ctx      context.Context
	Core     *appctx.AppCore // TODO: Maybe just rename to Env and make Utils a service.
	Services *AppServices
}

type AppServices struct {
	ProfileManager *profile.ProfileManager
	GameManager    *game.GameManager
	SteamLauncher  *steam.SteamLauncher
	TSAPI          *thunderstore.ThunderstoreAPI
	TSSchema       *thunderstore.ThunderstoreSchema
	TSDevTools     *thunderstore.ThunderstoreDevTools
}

func NewApplication() *Application {
	core := appctx.NewAppCore()
	services := NewAppServices(core)

	return &Application{
		Ctx:      nil,
		Core:     core,
		Services: services,
	}
}

func NewAppServices(core *appctx.AppCore) *AppServices {
	services := &AppServices{
		GameManager:    game.NewGameManager(),
		ProfileManager: profile.NewProfileManager(),
		SteamLauncher:  steam.NewSteamLauncher(core.Settings),
		TSAPI:          thunderstore.NewThunderstoreAPI(),
		TSSchema:       thunderstore.NewThunderstoreSchema(),
		TSDevTools:     thunderstore.NewThunderstoreDevTools(),
	}

	services.TSAPI.SetSchema(services.TSSchema)

	return services
}

func (app *Application) GetSettings() *appctx.AppSettings {
	return app.Core.Settings
}

func (app *Application) GetPersistence() *appctx.Persistence {
	return app.Core.Persistence
}

func (app *Application) GetUtils() *appctx.Utils {
	return app.Core.Utils
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
func (app *Application) Startup(ctx context.Context) {
	app.Ctx = ctx

	if app.GetPersistence().WindowState.Maximized {
		wuntime.WindowMaximise(ctx)
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
