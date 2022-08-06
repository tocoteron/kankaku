package model

import "github.com/tocoteron/kankaku/domain/model"

type Post struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	AuthorID string
}

func PostFrom(post *model.Post) *Post {
	return &Post{
		ID:       post.ID().String(),
		Content:  post.Content(),
		AuthorID: post.AuthorID().String(),
	}
}

func PostsFrom(posts *[]model.Post) []*Post {
	ps := []*Post{}
	for _, p := range *posts {
		ps = append(ps, PostFrom(&p))
	}
	return ps
}
