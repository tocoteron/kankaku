package service

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model"
	"github.com/tocoteron/kankaku/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (us *UserService) GetUser(id model.UserID) (*model.User, error) {
	u, err := us.repo.FindUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return u, nil
}

func (us *UserService) Post(id model.UserID, post model.Post) error {
	if err := us.repo.AddPost(id, post); err != nil {
		return fmt.Errorf("failed to post: %w", err)
	}
	return nil
}
