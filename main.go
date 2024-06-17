package main

import (
	"embed"
	"runtime"

	"github.com/leaanthony/u"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"modm8/backend"
	"modm8/backend/core"
	"modm8/backend/nexus"
	"modm8/backend/thunderstore"
)

//go:embed all:frontend/dist
var assets embed.FS

type IList []interface{}

func NewWindowsOptions(gpuAccel bool) *windows.Options {
	return &windows.Options{
		WindowIsTranslucent:  true,
		WebviewIsTransparent: true,
		BackdropType:         windows.Mica,
		ResizeDebounceMS:     1,
		WebviewUserDataPath:  core.ConfigDir(),
		WebviewGpuIsDisabled: !gpuAccel,
	}
}

func NewMacOptions() *mac.Options {
	return &mac.Options{
		TitleBar:             mac.TitleBarHiddenInset(),
		Appearance:           mac.NSAppearanceNameDarkAqua,
		WindowIsTranslucent:  true,
		WebviewIsTransparent: true,
		Preferences: &mac.Preferences{
			TabFocusesLinks: u.NewBool(false),
		},
	}
}

func main() {
	app := core.NewApp()
	app.Init()

	runtime.GOMAXPROCS(int(app.Settings.Performance.ThreadCount))

	nexusAPI := nexus.NewAPI()
	tsAPI := thunderstore.NewAPI()
	tsTools := thunderstore.NewTools()

	gameManager := backend.NewGameManager()

	err := wails.Run(&options.App{
		Title:     "modm8",
		Width:     int(app.Persistence.Window.Width),
		Height:    int(app.Persistence.Window.Height),
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:                true,
		EnableDefaultContextMenu: false,
		OnStartup:                app.Startup,
		OnBeforeClose:            app.OnBeforeClose,
		OnShutdown:               app.Shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "7465fe36-08e3-478b-853b-0f8676f724b7",
		},
		Mac:      NewMacOptions(),
		Windows:  NewWindowsOptions(app.Settings.Performance.GPUAcceleration),
		LogLevel: logger.INFO,
		Bind: IList{
			app,
			app.Settings,
			app.Persistence,
			tsAPI,
			tsTools,
			nexusAPI,
			gameManager,
		},
		EnumBind: IList{
			core.UpdateBehaviours,
			core.GameSelectionLayouts,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
