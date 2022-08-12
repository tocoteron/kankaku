package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tocoteron/kankaku/test"
)

func TestNewUser(t *testing.T) {
	id, _ := NewUserID("User ID")
	require.NotNil(t, id)

	tooLongName := test.RandomString(33)
	maxLenName := test.RandomString(32)

	tests := map[string]struct {
		id    UserID
		name  string
		posts []Post
		want  *User
	}{
		"empty name": {
			id:    *id,
			name:  "",
			posts: []Post{},
			want:  nil,
		},
		"too long name": {
			id:    *id,
			name:  tooLongName,
			posts: []Post{},
			want:  nil,
		},
		"max length name": {
			id:    *id,
			name:  maxLenName,
			posts: []Post{},
			want: &User{
				id:    *id,
				name:  maxLenName,
				posts: []Post{},
			},
		},
		"normal name": {
			id:    *id,
			name:  "Test User",
			posts: []Post{},
			want: &User{
				id:    *id,
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
