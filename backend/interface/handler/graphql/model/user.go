package model

import "github.com/tocoteron/kankaku/domain/model/user"

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

func UserFrom(u *user.User) *User {
	ps := []*Post{}
	for _, p := range u.Posts() {
		ps = append(ps, PostFrom(&p))
	}

	return &User{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Posts: ps,
	}
}
