package model

import "github.com/tocoteron/kankaku/domain/model/post"

type Post struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	AuthorID string
}

func PostFrom(post *post.Post) *Post {
	return &Post{
		ID:      post.ID().String(),
		Content: post.Content(),
	}
}

func PostsFrom(posts *[]post.Post) []*Post {
	ps := []*Post{}
	for _, p := range *posts {
		ps = append(ps, PostFrom(&p))
	}
	return ps
}
