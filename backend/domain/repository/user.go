package repository

import (
	"github.com/tocoteron/kankaku/domain/model"
)

type UserRepository interface {
	Save(user model.User) error
	Find(id model.UserID) (*model.User, error)
	GetAllPosts() (*[]model.Post, error)
	NextUserID() (*model.UserID, error)
	NextPostID() (*model.PostID, error)
}
