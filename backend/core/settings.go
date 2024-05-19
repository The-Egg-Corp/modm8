package core

import (
	"fmt"
	"os"
	"path"
	"runtime"

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
var AllUpdateBehaviours = []struct {
	Value  UpdateBehaviour
	TSName string
}{
	{UpdateBehaviourOff, "OFF"},
	{UpdateBehaviourNotify, "NOTIFY"},
	{UpdateBehaviourAuto, "AUTO"},
}

type PerformanceOptions struct {
	ThreadCount     uint8 `json:"thread_count" mapstructure:"thread_count"`
	GPUAcceleration bool  `json:"gpu_acceleration" mapstructure:"gpu_acceleration"`
}

type MiscOptions struct {
	AnimationsEnabled bool            `json:"animations_enabled" mapstructure:"animations_enabled"`
	UpdateBehaviour   UpdateBehaviour `json:"update_behaviour" mapstructure:"update_behaviour"`
}

type AppSettings struct {
	Locale      string             `json:"locale" mapstructure:"locale"`
	Theme       string             `json:"theme" mapstructure:"theme"`
	Performance PerformanceOptions `json:"performance" mapstructure:"performance"`
	Misc        MiscOptions        `json:"misc" mapstructure:"misc"`
}

// type AppSettingsService struct {
// 	Store    *viper.Viper `json:"store" mapstructure:"store"`
// 	Settings *AppSettings `json:"settings" mapstructure:"settings"`
// }

var cfg = viper.New()

func NewSettings() *AppSettings {
	return &AppSettings{
		Locale: "en",
		Theme:  "default-purple/dark",
		Performance: PerformanceOptions{
			ThreadCount:     uint8(runtime.NumCPU()),
			GPUAcceleration: true,
		},
		Misc: MiscOptions{
			AnimationsEnabled: true,
			UpdateBehaviour:   UpdateBehaviourAuto,
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

	if err := cfg.Unmarshal(settings); err != nil {
		return err
	}

	return nil
}

func (settings *AppSettings) Save() error {
	// Write struct values to viper config
	settings.WriteToConfig()

	// Write the config 'settings.toml' file.
	if err := cfg.WriteConfigAs(SettingsPath()); err != nil {
		return err
	}

	return nil
}

func (settings *AppSettings) SetLocale(locale string) {
	settings.Locale = locale
}

func (settings *AppSettings) SetTheme(theme string) {
	settings.Theme = theme
}

func (settings *AppSettings) SetThreads(count uint8) {
	settings.Performance.ThreadCount = count
}

func (settings *AppSettings) SetGPUAccel(val bool) {
	settings.Performance.GPUAcceleration = val
}

func (settings *AppSettings) SetAnimationsEnabled(val bool) {
	settings.Misc.AnimationsEnabled = val
}

func (settings *AppSettings) SetUpdateBehaviour(behaviour UpdateBehaviour) {
	settings.Misc.UpdateBehaviour = behaviour
}
