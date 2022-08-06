package model

import "github.com/tocoteron/kankaku/domain/model/post"

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func PostFrom(p *post.Post) *Post {
	return &Post{
		ID:      p.ID().String(),
		Content: p.Content(),
	}
}
