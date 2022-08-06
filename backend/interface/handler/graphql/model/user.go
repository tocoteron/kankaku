package model

import "github.com/tocoteron/kankaku/domain/model"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserFrom(u *model.User) *User {
	return &User{
		ID:   u.ID().String(),
		Name: u.Name(),
	}
}
