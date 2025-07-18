package appcore

import (
	"modm8/backend/common/paths"

	"github.com/spf13/viper"
)

var settingsCfg = viper.New()

type UpdateBehaviour uint8

const (
	UPDATE_BEHAVIOUR_OFF UpdateBehaviour = iota
	UPDATE_BEHAVIOUR_NOTIFY
	UPDATE_BEHAVIOUR_AUTO
)

type GameSelectionLayout string

const (
	GAME_SELECTION_LAYOUT_GRID GameSelectionLayout = "grid"
	GAME_SELECTION_LAYOUT_LIST GameSelectionLayout = "list"
)

type GeneralOptions struct {
	Locale            string          `json:"locale" mapstructure:"locale"`
	Theme             string          `json:"theme" mapstructure:"theme"`
	AnimationsEnabled bool            `json:"animations_enabled" mapstructure:"animations_enabled"`
	UpdateBehaviour   UpdateBehaviour `json:"update_behaviour" mapstructure:"update_behaviour"`
}

type PerformanceOptions struct {
	ThreadCount     uint8 `json:"thread_count" mapstructure:"thread_count"`
	GPUAcceleration bool  `json:"gpu_acceleration" mapstructure:"gpu_acceleration"`
}

type MiscOptions struct {
	SteamInstallPath    *string             `json:"steam_install_path" mapstructure:"steam_install_path"`
	NexusPersonalKey    *string             `json:"nexus_personal_key" mapstructure:"nexus_personal_key"`
	GameSelectionLayout GameSelectionLayout `json:"game_selection_layout" mapstructure:"game_selection_layout"`
}

type AppSettings struct {
	General     GeneralOptions     `json:"general" mapstructure:"general"`
	Performance PerformanceOptions `json:"performance" mapstructure:"performance"`
	Misc        MiscOptions        `json:"misc" mapstructure:"misc"`
}

func NewSettings() *AppSettings {
	return &AppSettings{
		General: GeneralOptions{
			Locale:            "en",
			Theme:             "aura-dark-purple",
			AnimationsEnabled: true,
			UpdateBehaviour:   UPDATE_BEHAVIOUR_AUTO,
		},
		Performance: PerformanceOptions{
			ThreadCount:     NumCPU(),
			GPUAcceleration: true,
		},
		Misc: MiscOptions{
			SteamInstallPath:    nil,
			NexusPersonalKey:    nil,
			GameSelectionLayout: "grid",
		},
	}
}

// func (settings *AppSettings) WriteToConfig() error {
// 	return WriteToConfig(settingsCfg, settings)
// }

func (settings *AppSettings) Load() error {
	SetupConfig(settingsCfg, "settings", "toml")
	return ReadOrCreate(settingsCfg, settings, paths.SettingsPath())
}

func (settings *AppSettings) Save() error {
	return Save(settingsCfg, settings, paths.SettingsPath())
}

func (settings *AppSettings) SaveAndApply() error {
	settings.Apply()
	return settings.Save()
}

// Any backend setting that can be updated without needing a restart should go here.
func (settings *AppSettings) Apply() {
	// Set the max number of processes (OS threads) to the value from the settings.toml file.
	SetMaxProcs(settings.Performance.ThreadCount)
}

func (settings *AppSettings) SetLocale(locale string) {
	settings.General.Locale = locale
}

func (settings *AppSettings) SetTheme(theme string) {
	settings.General.Theme = theme
}

func (settings *AppSettings) SetAnimationsEnabled(val bool) {
	settings.General.AnimationsEnabled = val
}

func (settings *AppSettings) SetUpdateBehaviour(behaviour UpdateBehaviour) {
	settings.General.UpdateBehaviour = behaviour
}

func (settings *AppSettings) SetThreads(count uint8) {
	settings.Performance.ThreadCount = count
}

func (settings *AppSettings) SetGPUAccel(val bool) {
	settings.Performance.GPUAcceleration = val
}

func (settings *AppSettings) SetSteamInstallPath(path string) {
	settings.Misc.SteamInstallPath = &path
}

func (settings *AppSettings) SetNexusPersonalKey(key string) {
	settings.Misc.NexusPersonalKey = &key
}

func (settings *AppSettings) SetGameSelectionLayout(layout GameSelectionLayout) {
	settings.Misc.GameSelectionLayout = layout
}
