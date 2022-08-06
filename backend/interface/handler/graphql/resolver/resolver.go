package resolver

import (
	"github.com/tocoteron/kankaku/interface/app"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *app.App
}

func NewResolver(app *app.App) *Resolver {
	return &Resolver{
		app: app,
	}
}
