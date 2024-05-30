package core

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type UpdateBehaviour uint8

const (
	UpdateBehaviourOff UpdateBehaviour = iota
	UpdateBehaviourNotify
	UpdateBehaviourAuto
)

// Mainly just for binding with Wails
var UpdateBehaviours = []struct {
	Value  UpdateBehaviour
	TSName string
}{
	{UpdateBehaviourOff, "OFF"},
	{UpdateBehaviourNotify, "NOTIFY"},
	{UpdateBehaviourAuto, "AUTO"},
}

type AppSettings struct {
	General     GeneralOptions     `json:"general" mapstructure:"general"`
	Performance PerformanceOptions `json:"performance" mapstructure:"performance"`
	Misc        MiscOptions        `json:"misc" mapstructure:"misc"`
	Video       VideoOptions       `json:"video" mapstructure:"video"`
}

type GeneralOptions struct {
	Locale            string `json:"locale" mapstructure:"locale"`
	Theme             string `json:"theme" mapstructure:"theme"`
	AnimationsEnabled bool   `json:"animations_enabled" mapstructure:"animations_enabled"`
}

type PerformanceOptions struct {
	ThreadCount     uint8 `json:"thread_count" mapstructure:"thread_count"`
	GPUAcceleration bool  `json:"gpu_acceleration" mapstructure:"gpu_acceleration"`
}

type MiscOptions struct {
	NexusPersonalKey    string          `json:"nexus_personal_key" mapstructure:"nexus_personal_key"`
	UpdateBehaviour     UpdateBehaviour `json:"update_behaviour" mapstructure:"update_behaviour"`
	GameSelectionLayout string          `json:"game_selection_layout" mapstructure:"game_selection_layout"`
}

type VideoOptions struct {
	WindowWidth  uint16 `json:"window_width" mapstructure:"window_width"`
	WindowHeight uint16 `json:"window_height" mapstructure:"window_height"`
	WindowPosX   uint32 `json:"window_pos_x" mapstructure:"window_pos_x"`
	WindowPosY   uint32 `json:"window_pos_y" mapstructure:"window_pos_y"`
}

var cfg = viper.New()

func NewSettings() *AppSettings {
	return &AppSettings{
		General: GeneralOptions{
			Locale:            "en",
			Theme:             "default-purple/dark",
			AnimationsEnabled: true,
		},
		Performance: PerformanceOptions{
			ThreadCount:     NumCPU(),
			GPUAcceleration: true,
		},
		Misc: MiscOptions{
			NexusPersonalKey:    "",
			UpdateBehaviour:     UpdateBehaviourAuto,
			GameSelectionLayout: "grid",
		},
		// Not accessible within settings, only in settings.toml for persistence.
		Video: VideoOptions{
			WindowWidth:  1380,
			WindowHeight: 930,
			WindowPosX:   0,
			WindowPosY:   0,
		},
	}
}

func ConfigDir() string {
	dir, _ := os.UserConfigDir()
	return path.Join(dir, "modm8")
}

func SettingsPath() string {
	return path.Join(ConfigDir(), "settings.toml")
}

func (settings *AppSettings) WriteToConfig() {
	var result map[string]interface{}
	err := mapstructure.Decode(settings, &result)
	if err != nil {
		fmt.Printf("%#v", result)
		return
	}

	cfg.MergeConfigMap(result)
}

func (settings *AppSettings) Load() error {
	cfg.SetConfigName("settings")
	cfg.SetConfigType("toml")
	cfg.AddConfigPath(ConfigDir())

	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return settings.Save()
		}

		return err
	}

	if err := cfg.UnmarshalExact(settings); err != nil {
		return err
	}

	return nil
}

func (settings *AppSettings) Save() error {
	// Write struct values to viper config
	settings.WriteToConfig()

	// Write config to the `settings.toml` file.
	if err := cfg.WriteConfigAs(SettingsPath()); err != nil {
		return err
	}

	return nil
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

func (settings *AppSettings) SetThreads(count uint8) {
	settings.Performance.ThreadCount = count
}

func (settings *AppSettings) SetGPUAccel(val bool) {
	settings.Performance.GPUAcceleration = val
}

func (settings *AppSettings) SetNexusPersonalKey(key string) {
	settings.Misc.NexusPersonalKey = key
}

func (settings *AppSettings) SetUpdateBehaviour(behaviour UpdateBehaviour) {
	settings.Misc.UpdateBehaviour = behaviour
}

func (settings *AppSettings) SetWindowWidth(width uint16) {
	settings.Video.WindowWidth = width
}

func (settings *AppSettings) SetWindowHeight(height uint16) {
	settings.Video.WindowHeight = height
}

func (settings *AppSettings) SetWindowX(xPos uint32) {
	settings.Video.WindowPosX = xPos
}

func (settings *AppSettings) SetWindowY(yPos uint32) {
	settings.Video.WindowPosY = yPos
}
