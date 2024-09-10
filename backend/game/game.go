package game

type PackageLoader int

const (
	ANCIENT_DUNGEON_VR = iota
	BEPINEX
	GODOT_ML
	LOVELY
	MELON_LOADER
	NORTHSTAR
	RETURN_OF_MODDING
	SHIMLOADER
)

type Game struct {
	DisplayName        string        `json:"display_name" mapstructure:"display_name"`
	SteamFolderName    string        `json:"steam_folder_name" mapstructure:"steam_folder_name"`
	SettingsIdentifier string        `json:"settings_identifier" mapstructure:"settings_identifier"`
	ExeName            []string      `json:"exe_name" mapstructure:"exe_name"`
	DataFolderName     string        `json:"data_folder_name" mapstructure:"data_folder_name"`
	ThunderstoreURL    *string       `json:"thunderstore_url" mapstructure:"thunderstore_url"`
	NexusURL           *string       `json:"nexus_url" mapstructure:"nexus_url"`
	ExclusionsURL      string        `json:"exclusions_url" mapstructure:"exclusions_url"`
	PackageLoader      PackageLoader `json:"package_loader" mapstructure:"package_loader"`
}
