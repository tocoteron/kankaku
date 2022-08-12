package model

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	id    UserID
	name  string
	posts []Post
}

func NewUser(id UserID, name string, posts []Post) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	u := &User{
		id:    id,
		name:  name,
		posts: posts,
	}

	return u, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Posts() []Post {
	return u.posts
}

func (u *User) Equals(other User) bool {
	return u.id.Equals(other.id)
}

func (u *User) Post(p Post) {
	u.posts = append(u.posts, p)
}

func validateName(name string) error {
	if utf8.RuneCountInString(name) == 0 {
		return fmt.Errorf("name must be not empty")
	}

	if utf8.RuneCountInString(name) > 32 {
		return fmt.Errorf("name must be 32 charactes or less")
	}

	return nil
}
