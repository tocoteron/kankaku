package model

import "github.com/tocoteron/kankaku/domain/model/user"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserFrom(u *user.User) *User {
	return &User{
		ID:   u.ID().String(),
		Name: u.Name(),
	}
}
