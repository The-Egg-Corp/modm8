package core

import (
	"context"

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

func (persistence *Persistence) WriteToConfig() error {
	return WriteToConfig(*persistenceCfg, persistence)
}

func (persistence *Persistence) Load() error {
	SetupConfig(*persistenceCfg, "persistence", "toml")
	return ReadOrCreate(*persistenceCfg, persistence)
}

func (persistence *Persistence) Save() error {
	return Save(*persistenceCfg, persistence)
}

// The frontend must still be loaded to call these runtime methods.
func (ws *Persistence) ApplyCurrentWindowState(ctx context.Context) {
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
