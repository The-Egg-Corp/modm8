package core

import (
	"context"

	"github.com/spf13/viper"

	wRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Persistence struct {
	Window         WindowState `json:"window" mapstructure:"window"`
	FavouriteGames []string    `json:"favourite_games" mapstructure:"favourite_games"`
}

type WindowState struct {
	Maximized bool   `json:"maximized" mapstructure:"maximized"`
	Width     uint16 `json:"width" mapstructure:"width"`
	Height    uint16 `json:"height" mapstructure:"height"`
	X         int    `json:"pos_x" mapstructure:"pos_x"`
	Y         int    `json:"pos_y" mapstructure:"pos_y"`
}

func NewPersistence() *Persistence {
	return &Persistence{
		FavouriteGames: []string{},
		Window: WindowState{
			Maximized: false,
			Width:     1380,
			Height:    930,
			X:         0,
			Y:         0,
		},
	}
}

func (persistence *Persistence) WriteToConfig() {
	WriteToConfig(*persistenceCfg, persistence)
}

func (persistence *Persistence) Load() error {
	SetupConfig(*persistenceCfg, "persistence", "toml")

	if err := persistenceCfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return persistence.Save()
		}

		return err
	}

	if err := persistenceCfg.Unmarshal(persistence); err != nil {
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

// To get the current size from the runtime, the frontend has to still be loaded.
func (ws *Persistence) SaveCurrentWindowState(ctx context.Context) {
	ws.SetMaximized(wRuntime.WindowIsMaximised(ctx))

	w, h := wRuntime.WindowGetSize(ctx)
	ws.SetWindowWidth(uint16(w))
	ws.SetWindowHeight(uint16(h))

	x, y := wRuntime.WindowGetPosition(ctx)
	ws.SetWindowX(x)
	ws.SetWindowY(y)
}

func (ws *Persistence) SetWindowWidth(width uint16) {
	ws.Window.Width = width
}

func (ws *Persistence) SetWindowHeight(height uint16) {
	ws.Window.Height = height
}

func (ws *Persistence) SetWindowX(xPos int) {
	ws.Window.X = xPos
}

func (ws *Persistence) SetWindowY(yPos int) {
	ws.Window.Y = yPos
}

func (ws *Persistence) SetMaximized(maximized bool) {
	ws.Window.Maximized = maximized
}

func (ws *Persistence) SetFavouriteGames(identifiers []string) {
	ws.FavouriteGames = identifiers
}
