package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tocoteron/kankaku/test"
)

func TestNewPost(t *testing.T) {
	id, _ := NewPostID("Post ID")
	require.NotNil(t, id)

	authorID, _ := NewUserID("Author ID")
	require.NotNil(t, authorID)

	tooLongContent := test.RandomString(257)
	maxLenContent := test.RandomString(256)

	tests := map[string]struct {
		id       PostID
		authorID UserID
		content  string
		want     *Post
	}{
		"empty content": {
			id:       *id,
			authorID: *authorID,
			content:  "",
			want:     nil,
		},
		"too long content": {
			id:       *id,
			authorID: *authorID,
			content:  tooLongContent,
			want:     nil,
		},
		"max length content": {
			id:       *id,
			authorID: *authorID,
			content:  maxLenContent,
			want: &Post{
				id:       *id,
				authorID: *authorID,
				content:  maxLenContent,
			},
		},
		"normal content": {
			id:       *id,
			authorID: *authorID,
			content:  "Hello",
			want: &Post{
				id:       *id,
				authorID: *authorID,
				content:  "Hello",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := NewPost(tt.id, tt.authorID, tt.content)
			assert.Equal(t, tt.want, got)
		})
	}
}
