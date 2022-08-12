package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tocoteron/kankaku/domain/model"
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
	id       string
	authorID string
	content  string
}

func NewUserInMemoryRepository() *userInMemoryRepository {
	return &userInMemoryRepository{
		users: map[string]*userDTO{},
	}
}

func (r *userInMemoryRepository) Save(user model.User) error {
	uid := user.ID().String()
	r.users[uid] = fromUser(&user)
	return nil
}

func (r *userInMemoryRepository) Find(id model.UserID) (*model.User, error) {
	uid := id.String()

	u, ok := r.users[uid]
	if !ok {
		return nil, fmt.Errorf("failed to find user (%s)", uid)
	}

	uu, err := u.toUser()
	if err != nil {
		return nil, fmt.Errorf("failed to map found User DTO to User model: %w", err)
	}

	return uu, nil
}

func (r *userInMemoryRepository) GetAllPosts() (*[]model.Post, error) {
	ps := []model.Post{}
	for _, u := range r.users {
		for _, p := range u.posts {
			post, err := p.toPost()
			if err != nil {
				return nil, fmt.Errorf("failed to map Post DTO to Post model: %w", err)
			}
			ps = append(ps, *post)
		}
	}

	return &ps, nil
}

func (r *userInMemoryRepository) NextUserID() (*model.UserID, error) {
	id, err := model.NewUserID(uuid.New().String())
	if err != nil {
		return nil, fmt.Errorf("failed to create user id: %w", err)
	}

	return id, nil
}

func (r *userInMemoryRepository) NextPostID() (*model.PostID, error) {
	id, err := model.NewPostID(uuid.New().String())
	if err != nil {
		return nil, fmt.Errorf("failed to create post id: %w", err)
	}
	return id, nil
}

// Map User DTO to User domain model
func (udto *userDTO) toUser() (*model.User, error) {
	id, err := model.NewUserID(udto.id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user id: %w", err)
	}

	posts := []model.Post{}
	for _, p := range udto.posts {
		pp, err := p.toPost()
		if err != nil {
			return nil, fmt.Errorf("failed to map Post DTO to Post model: %w", err)
		}
		posts = append(posts, *pp)
	}

	u, err := model.NewUser(*id, udto.name, posts)
	if err != nil {
		return nil, fmt.Errorf("failed to create User model from User DTO: %w", err)
	}

	return u, nil
}

// Map Post DTO to Post domain model
func (pdto *postDTO) toPost() (*model.Post, error) {
	id, err := model.NewPostID(pdto.id)
	if err != nil {
		return nil, fmt.Errorf("failed to create post id: %w", err)
	}

	authorID, err := model.NewUserID(pdto.authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to create author id: %w", err)
	}

	p, err := model.NewPost(*id, *authorID, pdto.content)
	if err != nil {
		return nil, fmt.Errorf("failed to create Post model from Post DTO: %w", err)
	}

	return p, nil
}

// Map User domain model to User DTO
func fromUser(user *model.User) *userDTO {
	ps := []postDTO{}
	for _, p := range user.Posts() {
		ps = append(ps, *fromPost(&p))
	}

	return &userDTO{
		id:    user.ID().String(),
		name:  user.Name(),
		posts: ps,
	}
}

// Map Post domain model to Post DTO
func fromPost(post *model.Post) *postDTO {
	return &postDTO{
		id:       post.ID().String(),
		authorID: post.AuthorID().String(),
		content:  post.Content(),
	}
}
