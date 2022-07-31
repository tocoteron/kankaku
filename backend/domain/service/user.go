package service

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
	"github.com/tocoteron/kankaku/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (us *UserService) GetUser(id user.UserID) (*user.User, error) {
	u, err := us.repo.FindUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return u, nil
}

func (us *UserService) Post(id user.UserID, post post.Post) error {
	if err := us.repo.AddPost(id, post); err != nil {
		return fmt.Errorf("failed to post: %w", err)
	}
	return nil
}
