package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"modm8/backend/core"
	"modm8/backend/thunderstore"
)

//go:embed all:frontend/dist
var assets embed.FS

type IList []interface{}

var windowsOpts = &windows.Options{
	WindowIsTranslucent:  true,
	WebviewIsTransparent: true,
	BackdropType:         windows.Mica,
	ResizeDebounceMS:     1,
	WebviewUserDataPath:  core.ConfigDir(),
}

func main() {
	app := core.NewApp()
	tsapi := thunderstore.NewAPI()

	err := wails.Run(&options.App{
		Title:     "modm8",
		Width:     1380,
		Height:    930,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:                true,
		EnableDefaultContextMenu: false,
		Windows:                  windowsOpts,
		OnStartup:                app.Startup,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "7465fe36-08e3-478b-853b-0f8676f724b7",
		},
		LogLevel: logger.INFO,
		Bind: IList{
			app,
			app.Settings,
			tsapi,
		},
		EnumBind: IList{
			core.AllUpdateBehaviours,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
