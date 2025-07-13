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
	"modm8/backend/app/appcore"
	"modm8/backend/common/paths"
	"modm8/backend/loaders"
	"modm8/backend/platform"
)

//go:embed all:frontend/dist
var assets embed.FS

type EnumBinding[T any] []struct {
	Value  T
	TSName string
}

var ModPlatforms = EnumBinding[platform.ModPlatform]{
	{platform.NEXUS, "NEXUS_MODS"},
	{platform.THUNDERSTORE, "THUNDERSTORE"},
}

var ModLoaders = EnumBinding[loaders.ModLoaderType]{
	{loaders.BEPINEX, "BEPINEX"},
	{loaders.MELON, "MELON"},
	{loaders.LOVELY, "LOVELY"},
}

var UpdateBehaviours = EnumBinding[appcore.UpdateBehaviour]{
	{appcore.UPDATE_BEHAVIOUR_OFF, "OFF"},
	{appcore.UPDATE_BEHAVIOUR_NOTIFY, "NOTIFY"},
	{appcore.UPDATE_BEHAVIOUR_AUTO, "AUTO"},
}

var GameSelectionLayouts = EnumBinding[appcore.GameSelectionLayout]{
	{appcore.GAME_SELECTION_LAYOUT_GRID, "GRID"},
	{appcore.GAME_SELECTION_LAYOUT_LIST, "LIST"},
}

func NewWindowsOptions(gpuAccel bool) *windows.Options {
	return &windows.Options{
		WindowIsTranslucent:  true,
		WebviewIsTransparent: true,
		BackdropType:         windows.Mica,
		ResizeDebounceMS:     1,
		WebviewUserDataPath:  paths.ConfigDir(),
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
	application := app.NewApplication()
	errs := application.Init()

	if len(errs) > 0 {
		for _, err := range errs {
			println("Error occurred:", err)
		}
	}

	core := application.GetCore()
	services := application.GetServices()

	// STRUCT BINDINGS
	bindings := []any{
		application,
		core.Settings,
		core.Persistence,
		core.Utils,
		services.GameManager,
		services.ProfileManager,
		services.SteamLauncher,
		services.TSSchema,
		services.TSAPI,
		//services.TSDevTools.PackageValidator,
	}

	// ENUM BINDINGS
	enumBindings := []any{
		UpdateBehaviours,
		GameSelectionLayouts,
		ModLoaders,
		ModPlatforms,
	}

	// For now, avoid binding Nexus stuff in GH Actions since key file wont exist.
	// if os.Getenv("GITHUB_ACTIONS") == "false" {
	// 	nexusAPI, keyErr := nexus.NewAPI(modm8.Ctx)
	// 	if keyErr != nil {
	// 		fmt.Printf("\nfailed to initialize nexus client:\n\n%v", keyErr)
	// 	}

	// 	bindings = append(bindings, nexusAPI)
	// }

	persistence := application.GetPersistence()
	settings := application.GetSettings()

	err := wails.Run(&options.App{
		Title:     "modm8",
		Width:     int(persistence.WindowState.Width),
		Height:    int(persistence.WindowState.Height),
		MinWidth:  640, // God's resolution.
		MinHeight: 480,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:                true,
		EnableDefaultContextMenu: false,
		OnStartup:                application.Startup,
		OnBeforeClose:            application.OnBeforeClose,
		OnShutdown:               application.Shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "7465fe36-08e3-478b-853b-0f8676f724b7",
		},
		Mac:      NewMacOptions(),
		Linux:    NewLinuxOptions(),
		Windows:  NewWindowsOptions(settings.Performance.GPUAcceleration),
		LogLevel: logger.INFO,
		Bind:     bindings,
		EnumBind: enumBindings,
	})

	if err != nil {
		println("Error launching app:", err)
	}
}
