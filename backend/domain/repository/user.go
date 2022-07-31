package repository

import (
	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
)

type UserRepository interface {
	FindUserByID(id user.UserID) (*user.User, error)
	AddPost(id user.UserID, post post.Post) error
}
