package backend

type Game struct {
	DisplayName        string
	SteamFolderName    string
	SettingsIdentifier string
	ExeName            []string
	DataFolderName     string
	ThunderstoreURL    string
	ExclusionsURL      string
}

type StorePlatform string

const (
	STEAM            StorePlatform = "Steam"
	STEAM_DIRECT     StorePlatform = "Steam " // Add a space so that there's no conflict in the PlatformInterceptor listing
	EPIC_GAMES_STORE StorePlatform = "Epic Games Store"
	OCULUS_STORE     StorePlatform = "Oculus Store"
	ORIGIN           StorePlatform = "Origin / EA Desktop"
	XBOX_GAME_PASS   StorePlatform = "Xbox Game Pass"
	OTHER            StorePlatform = "Other"
)
