package repository

import (
	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
)

type UserRepository interface {
	Save(user user.User) error
	FindUser(id user.UserID) (*user.User, error)
	AddPost(id user.UserID, post post.Post) error
	NextUserID() (*user.UserID, error)
	NextPostID() (*post.PostID, error)
}
