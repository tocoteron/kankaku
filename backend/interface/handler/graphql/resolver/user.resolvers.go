package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tocoteron/kankaku/domain/model"
	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
	"github.com/tocoteron/kankaku/interface/handler/graphql/generated"
	dto "github.com/tocoteron/kankaku/interface/handler/graphql/model"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*dto.User, error) {
	uc, err := mycontext.GetUserContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user context: %w", err)
	}

	return r.Query().User(ctx, uc.ID)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*dto.User, error) {
	id, err := model.NewUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user id: %w", err)
	}

	u, err := r.app.UserUseCase().GetUser(*id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return dto.UserFrom(u), nil
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *dto.User) ([]*dto.Post, error) {
	id, err := model.NewUserID(obj.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user id: %w", err)
	}

	ps, err := r.app.UserUseCase().GetUserPosts(*id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user posts: %w", err)
	}
	return dto.PostsFrom(ps), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
