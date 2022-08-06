package service

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model"
	"github.com/tocoteron/kankaku/domain/repository"
)

type UserService interface {
	Post(id model.UserID, post model.Post) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) Post(id model.UserID, post model.Post) error {
	if err := us.repository.AddPost(id, post); err != nil {
		return fmt.Errorf("failed to add post: %w", err)
	}
	return nil
}
