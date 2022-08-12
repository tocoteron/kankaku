package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tocoteron/kankaku/test"
)

func TestNewPostID(t *testing.T) {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	tooLongID := test.RandomStringWithCharset(257, charset)
	maxLenID := test.RandomStringWithCharset(256, charset)

	tests := map[string]struct {
		id   string
		want *PostID
	}{
		"empty id": {
			id:   "",
			want: nil,
		},
		"too long id": {
			id:   tooLongID,
			want: nil,
		},
		"max length id": {
			id: maxLenID,
			want: &PostID{
				value: maxLenID,
			},
		},
		"normal id": {
			id: "Post ID",
			want: &PostID{
				value: "Post ID",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := NewPostID(tt.id)
			assert.Equal(t, tt.want, got)
		})
	}
}
