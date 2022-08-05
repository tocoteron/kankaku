// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}