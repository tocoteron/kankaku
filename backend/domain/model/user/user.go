package user

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/post"
)

type User struct {
	id    UserID
	name  string
	posts []post.Post
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
		posts: []post.Post{},
	}

	return u, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Equals(other User) bool {
	return u.id.Equals(other.id)
}

func (u *User) Post(p post.Post) {
	u.posts = append(u.posts, p)
}
