package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tocoteron/kankaku/domain/model/post"
	"github.com/tocoteron/kankaku/domain/model/user"
)

type userInMemoryRepository struct {
	users map[string]*userDTO
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
		users: map[string]*userDTO{},
	}
}

func (r *userInMemoryRepository) Save(user user.User) error {
	uid := user.ID().String()

	ps := []postDTO{}
	for _, p := range user.Posts() {
		ps = append(ps, postDTO{
			id:      p.ID().String(),
			content: p.Content(),
		})
	}

	r.users[uid] = &userDTO{
		id:    user.ID().String(),
		name:  user.Name(),
		posts: ps,
	}

	return nil
}

func (r *userInMemoryRepository) FindUser(id user.UserID) (*user.User, error) {
	uid := id.String()

	u, ok := r.users[uid]
	if !ok {
		return nil, fmt.Errorf("failed to find user (%s)", uid)
	}

	uu, err := u.mapUser()
	if err != nil {
		return nil, fmt.Errorf("failed to map found User DTO to User model: %w", err)
	}

	return uu, nil
}

func (r *userInMemoryRepository) AddPost(id user.UserID, post post.Post) error {
	uid := id.String()

	u, ok := r.users[uid]
	if !ok {
		return fmt.Errorf("failed to find user (%s)", uid)
	}

	u.posts = append(u.posts, postDTO{
		id:      post.ID().String(),
		content: post.Content(),
	})

	return nil
}

func (r *userInMemoryRepository) GetAllPosts() (*[]post.Post, error) {
	ps := []post.Post{}
	for _, u := range r.users {
		for _, p := range u.posts {
			post, err := p.mapPost()
			if err != nil {
				return nil, fmt.Errorf("failed to map Post DTO to Post model: %w", err)
			}
			ps = append(ps, *post)
		}
	}

	return &ps, nil
}

func (r *userInMemoryRepository) NextUserID() (*user.UserID, error) {
	id := user.NewUserID(uuid.New().String())
	return &id, nil
}

func (r *userInMemoryRepository) NextPostID() (*post.PostID, error) {
	id := post.NewPostID(uuid.New().String())
	return &id, nil
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
