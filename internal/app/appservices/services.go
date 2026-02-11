package appservices

import (
	"modm8/internal/game"
	"modm8/internal/launchers/steam"
	"modm8/internal/profile"
	"modm8/internal/thunderstore"
)

type AppServices struct {
	ProfileManager *profile.ProfileManager
	GameManager    *game.GameManager
	SteamLauncher  *steam.SteamLauncher
	TSAPI          *thunderstore.ThunderstoreAPI
	TSSchema       *thunderstore.ThunderstoreSchema
	TSDevTools     *thunderstore.ThunderstoreDevTools
}

func New() *AppServices {
	services := &AppServices{
		GameManager:    game.NewGameManager(),
		ProfileManager: profile.NewProfileManager(),
		SteamLauncher:  steam.NewSteamLauncher(),
		TSAPI:          thunderstore.NewThunderstoreAPI(),
		TSSchema:       thunderstore.NewThunderstoreSchema(),
		TSDevTools:     thunderstore.NewThunderstoreDevTools(),
	}

	services.TSAPI.SetSchema(services.TSSchema)
	return services
}
