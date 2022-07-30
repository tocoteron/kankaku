package graph

import (
	"context"
	"fmt"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// Resolvers dependent user context which includes user id
type UserContext struct {
	ID uint64 `json:"id"`
}

type userContextKey struct{}

func SetUserContext(ctx context.Context, uctx *UserContext) context.Context {
	return context.WithValue(ctx, userContextKey{}, uctx)
}

func GetUserContext(ctx context.Context) (*UserContext, error) {
	userContext := ctx.Value(userContextKey{})
	if userContext == nil {
		return nil, fmt.Errorf("failed to retrieve user context")
	}

	uc, ok := userContext.(*UserContext)
	if !ok {
		return nil, fmt.Errorf("failed to convert to UserContext")
	}

	return uc, nil
}
