package app

import "github.com/spf13/viper"

type UpdateBehaviour uint8

const (
	UpdateBehaviourOff UpdateBehaviour = iota
	UpdateBehaviourNotify
	UpdateBehaviourAuto
)

var UpdateBehaviours = []struct {
	Value  UpdateBehaviour
	TSName string
}{
	{UpdateBehaviourOff, "OFF"},
	{UpdateBehaviourNotify, "NOTIFY"},
	{UpdateBehaviourAuto, "AUTO"},
}

type GameSelectionLayout string

const (
	GameSelectionLayoutGrid GameSelectionLayout = "grid"
	GameSelectionLayoutList GameSelectionLayout = "list"
)

var GameSelectionLayouts = []struct {
	Value  GameSelectionLayout
	TSName string
}{
	{GameSelectionLayoutGrid, "GRID"},
	{GameSelectionLayoutList, "LIST"},
}

type AppSettings struct {
	General     GeneralOptions     `json:"general" mapstructure:"general"`
	Performance PerformanceOptions `json:"performance" mapstructure:"performance"`
	Misc        MiscOptions        `json:"misc" mapstructure:"misc"`
}

type GeneralOptions struct {
	Locale            string  `json:"locale" mapstructure:"locale"`
	Theme             string  `json:"theme" mapstructure:"theme"`
	AnimationsEnabled bool    `json:"animations_enabled" mapstructure:"animations_enabled"`
	SteamInstallPath  *string `json:"steam_install_path" mapstructure:"steam_install_path"`
}

type PerformanceOptions struct {
	ThreadCount     uint8 `json:"thread_count" mapstructure:"thread_count"`
	GPUAcceleration bool  `json:"gpu_acceleration" mapstructure:"gpu_acceleration"`
}

type MiscOptions struct {
	NexusPersonalKey    *string             `json:"nexus_personal_key" mapstructure:"nexus_personal_key"`
	UpdateBehaviour     UpdateBehaviour     `json:"update_behaviour" mapstructure:"update_behaviour"`
	GameSelectionLayout GameSelectionLayout `json:"game_selection_layout" mapstructure:"game_selection_layout"`
}

var settingsCfg = viper.New()

func NewSettings() *AppSettings {
	return &AppSettings{
		General: GeneralOptions{
			Locale:            "en",
			Theme:             "aura-dark-purple",
			AnimationsEnabled: true,
			SteamInstallPath:  nil,
		},
		Performance: PerformanceOptions{
			ThreadCount:     NumCPU(),
			GPUAcceleration: true,
		},
		Misc: MiscOptions{
			NexusPersonalKey:    nil,
			UpdateBehaviour:     UpdateBehaviourAuto,
			GameSelectionLayout: "grid",
		},
	}
}

func (settings *AppSettings) Load() error {
	SetupConfig(settingsCfg, "settings", "toml")
	return ReadOrCreate(settingsCfg, settings, SettingsPath())
}

func (settings *AppSettings) Save() error {
	return Save(settingsCfg, settings, SettingsPath())
}

func (settings *AppSettings) WriteToConfig() error {
	return WriteToConfig(settingsCfg, settings)
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

func (settings *AppSettings) SetSteamInstallPath(path string) {
	settings.General.SteamInstallPath = &path
}

func (settings *AppSettings) SetThreads(count uint8) {
	settings.Performance.ThreadCount = count
}

func (settings *AppSettings) SetGPUAccel(val bool) {
	settings.Performance.GPUAcceleration = val
}

func (settings *AppSettings) SetNexusPersonalKey(key string) {
	settings.Misc.NexusPersonalKey = &key
}

func (settings *AppSettings) SetUpdateBehaviour(behaviour UpdateBehaviour) {
	settings.Misc.UpdateBehaviour = behaviour
}

func (settings *AppSettings) SetGameSelectionLayout(layout GameSelectionLayout) {
	settings.Misc.GameSelectionLayout = layout
}