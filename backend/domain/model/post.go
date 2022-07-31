package model

import "fmt"

type Post struct {
	id      PostID
	content string
}

type PostID struct {
	id string
}

func NewPost(id PostID, content string) (*Post, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("content must be not empty")
	}

	if len(content) > 256 {
		return nil, fmt.Errorf("content must be 32 charactes or less")
	}

	p := &Post{
		id:      id,
		content: content,
	}

	return p, nil
}
