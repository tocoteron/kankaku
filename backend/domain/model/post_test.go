package model

import (
	"testing"

	"github.com/tocoteron/kankaku/test"
)

func TestNewPost(t *testing.T) {
	tests := []struct {
		id       PostID
		authorID UserID
		content  string
		valid    bool
	}{
		{NewPostID("Post0"), NewUserID("User0"), "", false},                     // Empty content
		{NewPostID("Post1"), NewUserID("User1"), test.RandomString(257), false}, // Too long content
		{NewPostID("Post2"), NewUserID("User2"), test.RandomString(256), true},  // Max length content
		{NewPostID("Post3"), NewUserID("User3"), "Hello", true},                 // Normal content
	}

	for _, tt := range tests {
		got, err := NewPost(tt.id, tt.authorID, tt.content)

		if tt.valid {
			if err != nil {
				t.Errorf("failed to create new post: %s", err)
			}
			if !tt.id.Equals(got.ID()) {
				t.Errorf("post id '%s' doesn't equal to '%s'", tt.id, got.ID())
			}
			if !tt.authorID.Equals(got.authorID) {
				t.Errorf("author id '%s' doesn't equal to '%s'", tt.authorID, got.AuthorID())
			}
			if tt.content != got.Content() {
				t.Errorf("content '%s' doesn't equal to '%s'", tt.content, got.Content())
			}
		}

		if !tt.valid && got != nil {
			t.Errorf("invalid post was created: %+v", got)
		}
	}
}
