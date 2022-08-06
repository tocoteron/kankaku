package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tocoteron/kankaku/domain/model/user"
	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
	"github.com/tocoteron/kankaku/interface/handler/graphql/generated"
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

	return model.UserFrom(u), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	u, err := r.app.UserUseCase().GetUser(user.NewUserID(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return model.UserFrom(u), nil
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	ps, err := r.app.UserUseCase().GetUserPosts(user.NewUserID(obj.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return model.PostsFrom(ps), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
