package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tocoteron/kankaku/test"
)

func TestNewPost(t *testing.T) {
	tooLongContent := test.RandomString(257)
	maxLenContent := test.RandomString(256)

	tests := map[string]struct {
		id       PostID
		authorID UserID
		content  string
		want     *Post
	}{
		"empty content": {
			id:       NewPostID("Post0"),
			authorID: NewUserID("User0"),
			content:  "",
			want:     nil,
		},
		"too long content": {
			id:       NewPostID("Post1"),
			authorID: NewUserID("User1"),
			content:  tooLongContent,
			want:     nil,
		},
		"max length content": {
			id:       NewPostID("Post2"),
			authorID: NewUserID("User2"),
			content:  maxLenContent,
			want: &Post{
				id:       NewPostID("Post2"),
				authorID: NewUserID("User2"),
				content:  maxLenContent,
			},
		},
		"normal content": {
			id:       NewPostID("Post3"),
			authorID: NewUserID("User3"),
			content:  "Hello",
			want: &Post{
				id:       NewPostID("Post3"),
				authorID: NewUserID("User3"),
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
