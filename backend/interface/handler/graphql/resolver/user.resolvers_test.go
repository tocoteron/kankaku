package resolver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tocoteron/kankaku/domain/model"

	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
	dto "github.com/tocoteron/kankaku/interface/handler/graphql/model"
)

func TestMe(t *testing.T) {
	app, r := setupResolver()

	me, _ := app.UserUseCase().CreateUser("test user")
	require.NotNil(t, me)

	tests := map[string]struct {
		ctx  context.Context
		want *dto.User
	}{
		"empty context": {
			ctx:  context.Background(),
			want: nil,
		},
		"non-existent user": {
			ctx: mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{},
			),
			want: nil,
		},
		"me": {
			ctx: mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{ID: me.ID().String()},
			),
			want: dto.UserFrom(me),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := r.Query().Me(tt.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser(t *testing.T) {
	app, r := setupResolver()

	u, _ := app.UserUseCase().CreateUser("test user")
	require.NotNil(t, u)

	tests := map[string]struct {
		ctx    context.Context
		userID string
		want   *dto.User
	}{
		"non-existent user": {
			context.Background(),
			"",
			nil,
		},
		"existing user": {
			context.Background(),
			u.ID().String(),
			dto.UserFrom(u),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := r.Query().User(tt.ctx, tt.userID)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPosts(t *testing.T) {
	app, r := setupResolver()

	u, _ := app.UserUseCase().CreateUser("test user")
	require.NotNil(t, u)

	p, _ := app.UserUseCase().CreatePost(u.ID(), "test post")
	require.NotNil(t, p)

	tests := map[string]struct {
		ctx  context.Context
		obj  *dto.User
		want []*dto.Post
	}{
		"non-existent user": {
			ctx:  context.Background(),
			obj:  &dto.User{},
			want: nil,
		},
		"existing user": {
			ctx:  context.Background(),
			obj:  dto.UserFrom(u),
			want: dto.PostsFrom(&[]model.Post{*p}),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := r.User().Posts(tt.ctx, tt.obj)
			assert.Equal(t, tt.want, got)
		})
	}
}
