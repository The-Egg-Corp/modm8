package appcore

import (
	"context"
	"modm8/backend/common/paths"

	"github.com/spf13/viper"
	wuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Persistence struct {
	WindowState    WindowState `json:"window" mapstructure:"window"`
	FavouriteGames []string    `json:"favourite_games" mapstructure:"favourite_games"`
}

type WindowState struct {
	Maximized bool   `json:"maximized" mapstructure:"maximized"`
	Width     uint16 `json:"width" mapstructure:"width"`
	Height    uint16 `json:"height" mapstructure:"height"`
	X         int    `json:"pos_x" mapstructure:"pos_x"`
	Y         int    `json:"pos_y" mapstructure:"pos_y"`
}

var persistenceCfg = viper.New()

func NewPersistence() *Persistence {
	return &Persistence{
		FavouriteGames: []string{},
		WindowState: WindowState{
			Maximized: false,
			Width:     1380,
			Height:    930,
			X:         0,
			Y:         0,
		},
	}
}

func (persistence *Persistence) WriteToConfig() error {
	return WriteToConfig(persistenceCfg, persistence)
}

func (persistence *Persistence) Load() error {
	SetupConfig(persistenceCfg, "persistence", "toml")
	return ReadOrCreate(persistenceCfg, persistence, paths.PersistencePath())
}

func (persistence *Persistence) Save() error {
	return Save(persistenceCfg, persistence, paths.PersistencePath())
}

// The frontend must still be loaded to call these runtime methods.
func (persistence *Persistence) ApplyCurrentWindowState(ctx context.Context) {
	x, y := wuntime.WindowGetPosition(ctx)
	persistence.SetWindowX(x)
	persistence.SetWindowY(y)

	maximized := wuntime.WindowIsMaximised(ctx)
	persistence.SetMaximized(maximized)

	if !maximized {
		w, h := wuntime.WindowGetSize(ctx)
		persistence.SetWindowWidth(uint16(w))
		persistence.SetWindowHeight(uint16(h))
	}
}

func (persistence *Persistence) SetWindowWidth(width uint16) {
	persistence.WindowState.Width = width
}

func (persistence *Persistence) SetWindowHeight(height uint16) {
	persistence.WindowState.Height = height
}

func (persistence *Persistence) SetWindowX(xPos int) {
	persistence.WindowState.X = xPos
}

func (persistence *Persistence) SetWindowY(yPos int) {
	persistence.WindowState.Y = yPos
}

func (persistence *Persistence) SetMaximized(maximized bool) {
	persistence.WindowState.Maximized = maximized
}

func (persistence *Persistence) SetFavouriteGames(identifiers []string) {
	persistence.FavouriteGames = identifiers
}
