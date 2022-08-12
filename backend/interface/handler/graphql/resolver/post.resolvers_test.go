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

func TestPost(t *testing.T) {
	app, r := setupResolver()

	u, _ := app.UserUseCase().CreateUser("User")
	require.NotNil(t, u)

	pid, _ := model.NewPostID("Post ID")
	require.NotNil(t, pid)
	p, _ := model.NewPost(*pid, u.ID(), "Post")
	require.NotNil(t, p)

	tests := map[string]struct {
		ctx     context.Context
		content string
		want    *dto.Post
	}{
		"empty context": {
			ctx:     context.Background(),
			content: "Post",
			want:    nil,
		},
		"non-existent user": {
			ctx: mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{},
			),
			content: "Post",
			want:    nil,
		},
		"existing user": {
			ctx: mycontext.SetUserContext(
				context.Background(),
				&mycontext.UserContext{ID: u.ID().String()},
			),
			content: p.Content(),
			want:    dto.PostFrom(p),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := r.Mutation().Post(tt.ctx, tt.content)
			if tt.want != nil && got != nil {
				assert.Equal(t, tt.want.Content, got.Content)
				assert.Equal(t, tt.want.AuthorID, got.AuthorID)
			}
		})
	}
}
