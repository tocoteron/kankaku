package resolver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tocoteron/kankaku/domain/model"
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

	me, _ := app.UserUseCase().CreateUser("test user")
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
				&mycontext.UserContext{},
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

func TestUser(t *testing.T) {
	app, r := setup()

	u, _ := app.UserUseCase().CreateUser("test user")
	require.NotNil(t, u)

	tests := []struct {
		ctx    context.Context
		userID string
		want   *dto.User
	}{
		{
			// Can't get user because specified user id is invalid
			context.Background(),
			"",
			nil,
		},
		{
			// Can get user because specified user id is valid
			context.Background(),
			u.ID().String(),
			dto.UserFrom(u),
		},
	}

	for _, tt := range tests {
		got, _ := r.Query().User(tt.ctx, tt.userID)
		assert.Equal(t, tt.want, got)
	}
}

func TestPosts(t *testing.T) {
	app, r := setup()

	u, _ := app.UserUseCase().CreateUser("test user")
	require.NotNil(t, u)

	p, _ := app.UserUseCase().CreatePost(u.ID(), "test post")
	require.NotNil(t, p)

	tests := []struct {
		ctx  context.Context
		obj  *dto.User
		want []*dto.Post
	}{
		{
			// Can't get posts because specified user id is invalid
			context.Background(),
			&dto.User{},
			nil,
		},
		{
			// Can get posts because specified user id is valid
			context.Background(),
			dto.UserFrom(u),
			dto.PostsFrom(&[]model.Post{*p}),
		},
	}

	for _, tt := range tests {
		got, _ := r.User().Posts(tt.ctx, tt.obj)
		assert.Equal(t, tt.want, got)
	}
}
