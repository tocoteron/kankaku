package model

import (
	"fmt"
	"unicode/utf8"
)

type Post struct {
	id       PostID
	authorID UserID
	content  string
}

func NewPost(id PostID, authorID UserID, content string) (*Post, error) {
	if err := validateContent(content); err != nil {
		return nil, err
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

func validateContent(content string) error {
	if utf8.RuneCountInString(content) == 0 {
		return fmt.Errorf("content must be not empty")
	}

	if utf8.RuneCountInString(content) > 256 {
		return fmt.Errorf("content must be 256 charactes or less")
	}

	return nil
}
