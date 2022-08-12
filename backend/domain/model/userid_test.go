package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tocoteron/kankaku/test"
)

func TestNewUserID(t *testing.T) {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	tooLongID := test.RandomStringWithCharset(257, charset)
	maxLenID := test.RandomStringWithCharset(256, charset)

	tests := map[string]struct {
		id   string
		want *UserID
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
			want: &UserID{
				value: maxLenID,
			},
		},
		"normal id": {
			id: "User ID",
			want: &UserID{
				value: "User ID",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := NewUserID(tt.id)
			assert.Equal(t, tt.want, got)
		})
	}
}
