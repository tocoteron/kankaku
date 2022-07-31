package repository

import "github.com/tocoteron/kankaku/domain/model"

type UserRepository interface {
	FindUserByID(id model.UserID) (*model.User, error)
	AddPost(id model.UserID, post model.Post) error
}
