package game

type GameInfo struct {
	Dir             string    `json:"dir" mapstructure:"dir"`
	Name            string    `json:"name" mapstructure:"name"`
	Aliases         *[]string `json:"aliases" mapstructure:"aliases"`
	ThunderstoreURL *string   `json:"thunderstore_url" mapstructure:"thunderstore_url"`
	NexusURL        *string   `json:"nexus_url" mapstructure:"nexus_url"`
}

func NewGameInfo(dirPath, name string, aliases ...string) GameInfo {
	return GameInfo{
		Dir:     dirPath,
		Name:    name,
		Aliases: &aliases,
	}
}

type SteamGameInfo struct {
	GameInfo
	FolderName         string `json:"steam_folder_name" mapstructure:"steam_folder_name"`
	DataFolderName     string `json:"data_folder_name" mapstructure:"data_folder_name"`
	SettingsIdentifier string `json:"settings_identifier" mapstructure:"settings_identifier"`
	ExclusionsURL      string `json:"exclusions_url" mapstructure:"exclusions_url"`
}
