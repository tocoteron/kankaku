package service

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model"
	"github.com/tocoteron/kankaku/domain/repository"
)

type UserService interface {
	Post(user model.User, post model.Post) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) Post(user model.User, post model.Post) error {
	user.Post(post)
	if err := us.repository.Save(user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}
