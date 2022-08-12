package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tocoteron/kankaku/test"
)

func TestNewUser(t *testing.T) {
	tooLongName := test.RandomString(33)
	maxLenName := test.RandomString(32)

	tests := map[string]struct {
		id    UserID
		name  string
		posts []Post
		want  *User
	}{
		"empty name": {
			id:    NewUserID("User0"),
			name:  "",
			posts: []Post{},
			want:  nil,
		},
		"too long name": {
			id:    NewUserID("User1"),
			name:  tooLongName,
			posts: []Post{},
			want:  nil,
		},
		"max length name": {
			id:    NewUserID("User2"),
			name:  maxLenName,
			posts: []Post{},
			want: &User{
				id:    NewUserID("User2"),
				name:  maxLenName,
				posts: []Post{},
			},
		},
		"normal name": {
			id:    NewUserID("User3"),
			name:  "Test User",
			posts: []Post{},
			want: &User{
				id:    NewUserID("User3"),
				name:  "Test User",
				posts: []Post{},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := NewUser(tt.id, tt.name, tt.posts)
			assert.Equal(t, tt.want, got)
		})
	}
}
