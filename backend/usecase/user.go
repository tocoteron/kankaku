package usecase

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
	"github.com/tocoteron/kankaku/domain/repository"
	"github.com/tocoteron/kankaku/domain/service"
)

type UserUseCase interface {
	CreateUser(name string) (*user.User, error)
	GetUser(id user.UserID) (*user.User, error)
	CreatePost(id user.UserID, content string) (*post.Post, error)
}

type userUseCase struct {
	service    service.UserService
	repository repository.UserRepository
}

func NewUserUseCase(service service.UserService, repository repository.UserRepository) *userUseCase {
	return &userUseCase{
		service:    service,
		repository: repository,
	}
}

func (u *userUseCase) CreateUser(name string) (*user.User, error) {
	id, err := u.repository.NextUserID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %w", err)
	}

	user, err := user.NewUser(*id, name, []post.Post{})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if err := u.repository.Save(*user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (u *userUseCase) GetUser(id user.UserID) (*user.User, error) {
	user, err := u.repository.FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (u *userUseCase) CreatePost(id user.UserID, content string) (*post.Post, error) {
	postID, err := u.repository.NextPostID()
	if err != nil {
		return nil, fmt.Errorf("failed to get next post id: %w", err)
	}

	p, err := post.NewPost(*postID, content)
	if err != nil {
		return nil, fmt.Errorf("failed to create new post: %w", err)
	}

	err = u.service.Post(id, *p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}
