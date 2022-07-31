package model

import (
	"fmt"
)

type User struct {
	id    UserID
	name  string
	posts []Post
}

type UserID struct {
	id string
}

func NewUser(id UserID, name string) (*User, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name must be not empty")
	}

	if len(name) > 32 {
		return nil, fmt.Errorf("name must be 32 charactes or less")
	}

	u := &User{
		id:    id,
		name:  name,
		posts: []Post{},
	}

	return u, nil
}

func (u *User) Post(p Post) {
	if u.id == u.id {

	}
	u.posts = append(u.posts, p)
}
