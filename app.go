package main

import (
	"context"
	"strings"

	"github.com/samber/lo"
	v1 "github.com/the-egg-corp/thundergo/v1"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetUserPackages(owner string) string {
	pkgs, err := v1.GetAllPackages()
	if err != nil {
		return "An error getting packages!"
	}

	pkgs = lo.Filter(pkgs, func(pkg v1.Package, index int) bool {
		return strings.EqualFold(pkg.Owner, owner)
	})

	if pkgs.Size() == 1 {
		return pkgs[0].Name
	}

	var names []string
	var i uint

	for i = 0; i < pkgs.Size(); i++ {
		names = append(names, pkgs[i].Name)
	}

	return strings.Join(names, ", ")
}
