package resolver

import (
	"github.com/tocoteron/kankaku/interface/app"
)

func setupResolver() (*app.App, *Resolver) {
	app := app.NewTestApp()
	resolver := NewResolver(app)

	return app, resolver
}
