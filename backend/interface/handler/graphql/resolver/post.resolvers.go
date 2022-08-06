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

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context, content string) (*model.Post, error) {
	uc, err := mycontext.GetUserContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user context: %w", err)
	}

	p, err := r.app.UserUseCase().CreatePost(user.NewUserID(uc.ID), content)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	post := &model.Post{
		ID:      p.ID().String(),
		Content: p.Content(),
	}

	return post, nil
}
