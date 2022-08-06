package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/user"
	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
	"github.com/tocoteron/kankaku/interface/handler/graphql/model"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	uc, err := mycontext.GetUserContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user context: %w", err)
	}

	u, err := r.app.UserUseCase().GetUser(user.NewUserID(uc.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	posts := []*model.Post{}
	for _, p := range u.Posts() {
		posts = append(posts, &model.Post{
			ID:      p.ID().String(),
			Content: p.Content(),
		})
	}

	user := &model.User{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Posts: posts,
	}

	return user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	u, err := r.app.UserUseCase().GetUser(user.NewUserID(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	posts := []*model.Post{}
	for _, p := range u.Posts() {
		posts = append(posts, &model.Post{
			ID:      p.ID().String(),
			Content: p.Content(),
		})
	}

	user := &model.User{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Posts: posts,
	}

	return user, nil
}
