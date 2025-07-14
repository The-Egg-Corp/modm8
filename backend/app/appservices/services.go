package appservices

import (
	"modm8/backend/app/appcore"
	"modm8/backend/game"
	"modm8/backend/launchers/steam"
	"modm8/backend/profile"
	"modm8/backend/thunderstore"
)

type AppServices struct {
	ProfileManager *profile.ProfileManager
	GameManager    *game.GameManager
	SteamLauncher  *steam.SteamLauncher
	TSAPI          *thunderstore.ThunderstoreAPI
	TSSchema       *thunderstore.ThunderstoreSchema
	TSDevTools     *thunderstore.ThunderstoreDevTools
}

func New(core *appcore.AppCore) *AppServices {
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
