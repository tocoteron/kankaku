package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
)

type userInMemoryRepository struct {
	users []userDTO
}

type userDTO struct {
	id    string
	name  string
	posts []postDTO
}

type postDTO struct {
	id      string
	content string
}

func NewUserInMemoryRepository() *userInMemoryRepository {
	return &userInMemoryRepository{
		users: []userDTO{},
	}
}

func (r *userInMemoryRepository) FindUser(id user.UserID) (*user.User, error) {
	uid := id.String()

	u, err := r.findUserByID(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to find user (%s): %w", uid, err)
	}

	uu, err := u.mapUser()
	if err != nil {
		return nil, fmt.Errorf("failed to map found User DTO to User model: %w", err)
	}

	return uu, nil
}

func (r *userInMemoryRepository) AddPost(id user.UserID, post post.Post) error {
	uid := id.String()

	u, err := r.findUserByID(uid)
	if err != nil {
		return fmt.Errorf("failed to find user (%s): %w", uid, err)
	}

	u.posts = append(u.posts, postDTO{
		id:      post.ID().String(),
		content: post.Content(),
	})

	return nil
}

func (r *userInMemoryRepository) NextUserID() (*user.UserID, error) {
	id := user.NewUserID(uuid.New().String())
	return &id, nil
}

func (r *userInMemoryRepository) NextPostID() (*post.PostID, error) {
	id := post.NewPostID(uuid.New().String())
	return &id, nil
}

// Helper function to find UserDTO from in-memory DB
func (r *userInMemoryRepository) findUserByID(id string) (*userDTO, error) {
	for _, u := range r.users {
		if u.id == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("failed to find user (%s) from in-memory DB", id)
}

// Map User DTO to User domain model
func (udto *userDTO) mapUser() (*user.User, error) {
	id := user.NewUserID(udto.id)

	posts := []post.Post{}
	for _, p := range udto.posts {
		pp, err := p.mapPost()
		if err != nil {
			return nil, fmt.Errorf("failed to map Post DTO to Post model: %w", err)
		}
		posts = append(posts, *pp)
	}

	u, err := user.NewUser(id, udto.name, posts)
	if err != nil {
		return nil, fmt.Errorf("failed to create User model from User DTO: %w", err)
	}

	return u, nil
}

// Map Post DTO to Post domain model
func (pdto *postDTO) mapPost() (*post.Post, error) {
	id := post.NewPostID(pdto.id)
	p, err := post.NewPost(id, pdto.content)
	if err != nil {
		return nil, fmt.Errorf("failed to create Post model from Post DTO: %w", err)
	}
	return p, nil
}
