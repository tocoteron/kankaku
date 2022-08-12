package resolver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tocoteron/kankaku/interface/app"

	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
	dto "github.com/tocoteron/kankaku/interface/handler/graphql/model"
)

func setup() (*app.App, *Resolver) {
	app := app.NewTestApp()
	resolver := NewResolver(app)

	return app, resolver
}

func TestMe(t *testing.T) {
	app, r := setup()
	me, _ := app.UserUseCase().CreateUser("test")
	require.NotNil(t, me)

	tests := []struct {
		ctx  context.Context
		want *dto.User
	}{
		{
			// Can't get user because context is empty
			context.Background(),
			nil,
		},
		{
			// Can't get user because specified user id is invalid
			mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{ID: "0"},
			),
			nil,
		},
		{
			// Can get user because specified user id is valid
			mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{ID: me.ID().String()},
			),
			dto.UserFrom(me),
		},
	}

	for _, tt := range tests {
		got, _ := r.Query().Me(tt.ctx)
		assert.Equal(t, tt.want, got)
	}
}
