package repository

import (
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
)

type UserInMemoryRepository struct {
	users []UserDTO
}

type UserDTO struct {
	id    string
	name  string
	posts []PostDTO
}

type PostDTO struct {
	id      string
	content string
}

func NewUserRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		users: []UserDTO{},
	}
}

func (r *UserInMemoryRepository) FindUserByID(id user.UserID) (*user.User, error) {
	uid := id.String()

	u, err := r.findUserByID(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	uu, err := u.mapUser()
	if err != nil {
		return nil, fmt.Errorf("failed to map found User DTO to User model: %w", err)
	}

	return uu, nil
}

func (r *UserInMemoryRepository) AddPost(id user.UserID, post post.Post) error {
	uid := id.String()

	u, err := r.findUserByID(uid)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	u.posts = append(u.posts, PostDTO{
		id:      post.ID().String(),
		content: post.Content(),
	})

	return nil
}

// Helper function to find UserDTO from in-memory DB
func (r *UserInMemoryRepository) findUserByID(id string) (*UserDTO, error) {
	for _, u := range r.users {
		if u.id == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("failed to find user from in-memory DB")
}

// Map User DTO to User domain model
func (udto *UserDTO) mapUser() (*user.User, error) {
	id := user.NewUserID(udto.id)

	posts := []post.Post{}
	for _, p := range udto.posts {
		pp, err := p.mapPost()
		if err != nil {
			return nil, fmt.Errorf("failed to map User DTO to User model: %w", err)
		}
		posts = append(posts, *pp)
	}

	u, err := user.NewUser(id, udto.name, posts)
	if err != nil {
		return nil, fmt.Errorf("failed to map User DTO to User model: %w", err)
	}

	return u, nil
}

// Map Post DTO to Post domain model
func (pdto *PostDTO) mapPost() (*post.Post, error) {
	id := post.NewPostID(pdto.id)
	p, err := post.NewPost(id, pdto.content)
	if err != nil {
		return nil, fmt.Errorf("failed to map Post DTO to Post model: %w", err)
	}
	return p, nil
}
