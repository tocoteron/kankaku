package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tocoteron/kankaku/graph/generated"
	"github.com/tocoteron/kankaku/graph/model"
)

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context, content string) (*model.Post, error) {
	return nil, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	uc, err := GetUserContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user context: %w", err)
	}

	id := strconv.FormatUint(uc.ID, 10)

	user := &model.User{
		ID:   id,
		Name: "User " + id,
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
