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

	return model.PostFrom(p), nil
}

// Author is the resolver for the author field.
func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	return r.Resolver.Query().User(ctx, obj.AuthorID)
}

// Timeline is the resolver for the timeline field.
func (r *queryResolver) Timeline(ctx context.Context) ([]*model.Post, error) {
	tl, err := r.app.UserUseCase().GetTimeline()
	if err != nil {
		return nil, fmt.Errorf("failed to get timeline: %w", err)
	}

	return model.PostsFrom(tl), nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
