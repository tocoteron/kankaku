package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tocoteron/kankaku/test"
)

func TestNewPost(t *testing.T) {
	tooLongContent := test.RandomString(257)
	maxLenContent := test.RandomString(256)

	tests := []struct {
		id       PostID
		authorID UserID
		content  string
		want     *Post
	}{
		{ // Empty content
			NewPostID("Post0"),
			NewUserID("User0"),
			"",
			nil,
		},
		{ // Too long content
			NewPostID("Post1"),
			NewUserID("User1"),
			tooLongContent,
			nil,
		},
		{ // Max length content
			NewPostID("Post2"),
			NewUserID("User2"),
			maxLenContent,
			&Post{
				id:       NewPostID("Post2"),
				authorID: NewUserID("User2"),
				content:  maxLenContent,
			},
		},
		{ // Normal content
			NewPostID("Post3"),
			NewUserID("User3"),
			"Hello",
			&Post{
				id:       NewPostID("Post3"),
				authorID: NewUserID("User3"),
				content:  "Hello",
			},
		},
	}

	for _, tt := range tests {
		got, _ := NewPost(tt.id, tt.authorID, tt.content)
		assert.Equal(t, tt.want, got)
	}
}
