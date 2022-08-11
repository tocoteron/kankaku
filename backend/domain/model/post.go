package model

import "fmt"

type Post struct {
	id       PostID
	authorID UserID
	content  string
}

func NewPost(id PostID, authorID UserID, content string) (*Post, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("content must be not empty")
	}

	if len(content) > 256 {
		return nil, fmt.Errorf("content must be 256 charactes or less")
	}

	p := &Post{
		id:       id,
		authorID: authorID,
		content:  content,
	}

	return p, nil
}

func (p *Post) ID() PostID {
	return p.id
}

func (p *Post) AuthorID() UserID {
	return p.authorID
}

func (p *Post) Content() string {
	return p.content
}

func (p *Post) Equals(other Post) bool {
	return p.id.Equals(other.id)
}
