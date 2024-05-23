package backend

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

type Game struct {
	DisplayName        string   `json:"display_name" mapstructure:"display_name"`
	SteamFolderName    string   `json:"steam_folder_name" mapstructure:"steam_folder_name"`
	SettingsIdentifier string   `json:"settings_identifier" mapstructure:"settings_identifier"`
	ExeName            []string `json:"exe_name" mapstructure:"exe_name"`
	DataFolderName     string   `json:"data_folder_name" mapstructure:"data_folder_name"`
	ThunderstoreURL    string   `json:"thunderstore_url" mapstructure:"thunderstore_url"`
	ExclusionsURL      string   `json:"exclusions_url" mapstructure:"exclusions_url"`
}
