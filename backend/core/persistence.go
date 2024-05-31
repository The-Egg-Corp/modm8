package core

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Persistence struct {
	Window WindowState
}

type WindowState struct {
	Width  uint16 `json:"width" mapstructure:"width"`
	Height uint16 `json:"height" mapstructure:"height"`
	X      uint32 `json:"pos_x" mapstructure:"pos_x"`
	Y      uint32 `json:"pos_y" mapstructure:"pos_y"`
}

var persistenceCfg = viper.New()

func NewPersistence() *Persistence {
	return &Persistence{
		Window: WindowState{
			Width:  1380,
			Height: 930,
			X:      0,
			Y:      0,
		},
	}
}

func (settings *Persistence) WriteToConfig() {
	var result map[string]interface{}
	err := mapstructure.Decode(settings, &result)
	if err != nil {
		fmt.Printf("%#v", result)
		return
	}

	persistenceCfg.MergeConfigMap(result)
}

func (settings *Persistence) Load() error {
	persistenceCfg.SetConfigName("persistence")
	persistenceCfg.SetConfigType("toml")
	persistenceCfg.AddConfigPath(ConfigDir())

	if err := persistenceCfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return settings.Save()
		}

		return err
	}

	if err := persistenceCfg.Unmarshal(settings); err != nil {
		return err
	}

	return nil
}

func (persistence *Persistence) Save() error {
	// Write struct values to viper config
	persistence.WriteToConfig()

	// Write config to the `persistence.toml` file.
	if err := persistenceCfg.WriteConfigAs(PersistencePath()); err != nil {
		return err
	}

	return nil
}

func (ws *Persistence) SetWindowWidth(width uint16) {
	ws.Window.Width = width
}

func (ws *Persistence) SetWindowHeight(height uint16) {
	ws.Window.Height = height
}

func (ws *Persistence) SetWindowX(xPos uint32) {
	ws.Window.X = xPos
}

func (ws *Persistence) SetWindowY(yPos uint32) {
	ws.Window.Y = yPos
}
