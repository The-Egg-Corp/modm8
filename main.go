package main

import (
	"embed"

	"github.com/leaanthony/u"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"modm8/backend/app"
	"modm8/backend/game"
	"modm8/backend/nexus"
	"modm8/backend/thunderstore"

	"modm8/backend/runners/steam"
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
		WebviewUserDataPath:  app.ConfigDir(),
		WebviewGpuIsDisabled: !gpuAccel,
	}
}

func NewLinuxOptions() *linux.Options {
	return &linux.Options{
		WindowIsTranslucent: true,
		WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
		ProgramName:         "modm8",
	}
}

func NewMacOptions() *mac.Options {
	return &mac.Options{
		TitleBar:             mac.TitleBarHiddenInset(),
		Appearance:           mac.NSAppearanceNameDarkAqua,
		WindowIsTranslucent:  true,
		WebviewIsTransparent: true,
		Preferences: &mac.Preferences{
			TabFocusesLinks: u.False,
		},
	}
}

func main() {
	modm8 := app.NewApp()
	errs := modm8.Init()

	if len(errs) > 0 {
		for _, err := range errs {
			println("Error occurred:", err)
		}
	}

	modm8.Settings.Apply()

	nexusAPI := nexus.NewAPI(modm8.Ctx)
	tsAPI := thunderstore.NewAPI(modm8.Ctx)
	tsTools := thunderstore.NewTools()

	gameManager := game.NewManager()
	steamRunner := steam.NewRunner()

	err := wails.Run(&options.App{
		Title:     "modm8",
		Width:     int(modm8.Persistence.Window.Width),
		Height:    int(modm8.Persistence.Window.Height),
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:                true,
		EnableDefaultContextMenu: false,
		OnStartup:                modm8.Startup,
		OnBeforeClose:            modm8.OnBeforeClose,
		OnShutdown:               modm8.Shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "7465fe36-08e3-478b-853b-0f8676f724b7",
		},
		Mac:      NewMacOptions(),
		Linux:    NewLinuxOptions(),
		Windows:  NewWindowsOptions(modm8.Settings.Performance.GPUAcceleration),
		LogLevel: logger.INFO,
		Bind: IList{
			modm8,
			modm8.Settings,
			modm8.Persistence,
			modm8.Utils,
			tsAPI,
			tsTools,
			nexusAPI,
			gameManager,
			steamRunner,
		},
		EnumBind: IList{
			app.UpdateBehaviours,
			app.GameSelectionLayouts,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
