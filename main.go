package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/leaanthony/u"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"modm8/backend/app"
	"modm8/backend/common/profile"
	"modm8/backend/game"
	"modm8/backend/loaders"
	"modm8/backend/nexus"
	"modm8/backend/thunderstore"

	"modm8/backend/runners/steam"
)

//go:embed all:frontend/dist
var assets embed.FS

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

// TESTING WITH ADMIN PRIVILEGES
// var showCmd int32 = 1 // SW_NORMAL

// // (Windows) Requests admin rights by prompting the user with a UAC dialog.
// // This function will relaunch the application with elevated privileges if the user accepts.
// func requestElevation(exe string) error {
// 	cwd, _ := os.Getwd()
// 	args := strings.Join(os.Args[1:], " ")

// 	verbPtr, _ := syscall.UTF16PtrFromString("runas")
// 	exePtr, _ := syscall.UTF16PtrFromString(exe)
// 	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
// 	argPtr, _ := syscall.UTF16PtrFromString(args)

// 	err := win.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
// 	if err != nil {
// 		fmt.Printf("failed to elevate privileges.\n%v\n", err)
// 	}

// 	return err
// }

// func main() {
// 	// Check and prompt user for admin rights on Windows.
// 	// This is required for symlink creation which we use when dealing with mod profiles.
// 	if runtime.GOOS == "windows" {
// 		// Lil hack so we don't get a UAC prompt on processes other than the final app.
// 		// At the very least, this fixes a wails-binding prompt in `wails dev`.
// 		exe, _ := os.Executable()
// 		correctProc := filepath.Base(exe) == "modm8.exe"

// 		if correctProc && !win.GetCurrentProcessToken().IsElevated() {
// 			requestElevation(exe)
// 			os.Exit(0)
// 		}
// 	}

// 	// Linux allows symlinks without admin. (based)
// 	launch()
// }

func main() {
	modm8 := app.NewApp()
	errs := modm8.Init()

	if len(errs) > 0 {
		for _, err := range errs {
			println("Error occurred:", err)
		}
	}

	modm8.Settings.Apply()

	tsAPI := thunderstore.NewThunderstoreAPI(modm8.Ctx)
	tsSchema := thunderstore.NewThunderstoreSchema()
	tsTools := thunderstore.NewThunderstoreTools()

	profileManager := profile.NewProfileManager()
	gameManager := game.NewGameManager()
	steamRunner := steam.NewSteamRunner()

	// STRUCT BINDINGS
	bindings := []any{
		modm8,
		modm8.Settings,
		modm8.Persistence,
		modm8.Utils,
		profileManager,
		gameManager,
		steamRunner,
		tsAPI,
		tsSchema,
		tsTools,
	}

	// ENUM BINDINGS
	enumBindings := []any{
		app.UpdateBehaviours,
		app.GameSelectionLayouts,
		loaders.ModLoaders,
	}

	// For now, avoid binding Nexus stuff in GH Actions since key file wont exist.
	if os.Getenv("GITHUB_ACTIONS") == "false" {
		nexusAPI, keyErr := nexus.NewAPI(modm8.Ctx)
		if keyErr != nil {
			fmt.Printf("\nfailed to initialize nexus client:\n\n%v", keyErr)
		}

		bindings = append(bindings, nexusAPI)
	}

	err := wails.Run(&options.App{
		Title:     "modm8",
		Width:     int(modm8.Persistence.Window.Width),
		Height:    int(modm8.Persistence.Window.Height),
		MinWidth:  640, // God's resolution.
		MinHeight: 480,
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
		Bind:     bindings,
		EnumBind: enumBindings,
	})

	if err != nil {
		println("Error launching app:", err)
	}
}
