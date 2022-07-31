package post

import "fmt"

type Post struct {
	id      PostID
	content string
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

func (p *Post) Equals(other Post) bool {
	return p.EqualsID(other.id)
}

func (p *Post) EqualsID(id PostID) bool {
	return p.id.Equals(id)
}
