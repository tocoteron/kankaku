package usecase

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model"
	"github.com/tocoteron/kankaku/domain/repository"
	"github.com/tocoteron/kankaku/domain/service"
)

type UserUseCase interface {
	CreateUser(name string) (*model.User, error)
	GetUser(id model.UserID) (*model.User, error)
	GetUserPosts(id model.UserID) (*[]model.Post, error)
	CreatePost(id model.UserID, content string) (*model.Post, error)
	GetTimeline() (*[]model.Post, error)
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

func (u *userUseCase) CreateUser(name string) (*model.User, error) {
	id, err := u.repository.NextUserID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %w", err)
	}

	user, err := model.NewUser(*id, name, []model.Post{})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if err := u.repository.Save(*user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (u *userUseCase) GetUser(id model.UserID) (*model.User, error) {
	user, err := u.repository.Find(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (u *userUseCase) GetUserPosts(id model.UserID) (*[]model.Post, error) {
	user, err := u.repository.Find(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	ps := user.Posts()
	return &ps, nil
}

func (u *userUseCase) CreatePost(id model.UserID, content string) (*model.Post, error) {
	postID, err := u.repository.NextPostID()
	if err != nil {
		return nil, fmt.Errorf("failed to get next post id: %w", err)
	}

	p, err := model.NewPost(*postID, id, content)
	if err != nil {
		return nil, fmt.Errorf("failed to create new post: %w", err)
	}

	err = u.service.Post(id, *p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}

func (u *userUseCase) GetTimeline() (*[]model.Post, error) {
	ps, err := u.repository.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to get all posts: %w", err)
	}

	return ps, nil
}
