package context

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/tocoteron/kankaku/auth"
	"github.com/tocoteron/kankaku/graph"
)

var EchoContextKey = &ContextKey{"EchoContextKey"}

type ContextKey struct {
	name string
}

// Retrieve echo.Context from context.Context
func EchoContextFromContext(ctx context.Context) (echo.Context, error) {
	echoContext := ctx.Value(EchoContextKey)
	if echoContext == nil {
		return nil, fmt.Errorf("could not retrieve echo.Context")
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		return nil, fmt.Errorf("echo.Context has wrong type")
	}
	return ec, nil
}

// Bind echo.Context to context.Context of http.Request
func ContextProvider() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), EchoContextKey, c)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

// Bind user context to context.Context of http.Request
func UserContextProvider() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Retrieve token from context
			jwtCtx := c.Get("user")
			if jwtCtx == nil {
				return fmt.Errorf("failed to retrieve token from context")
			}
			token, ok := jwtCtx.(*jwt.Token)
			if !ok {
				return fmt.Errorf("failed to convert token to *jwt.Token")
			}
			claims, ok := token.Claims.(*auth.JWTCustomClaims)
			if !ok {
				return fmt.Errorf("failed to convert jwt claims to *auth.JWTCustomClaims")
			}

			// Parse user id of token
			id, err := strconv.ParseUint(claims.Id, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse token id: %w", err)
			}

			// Set user context
			ctx := graph.SetUserContext(c.Request().Context(), &graph.UserContext{ID: id})
			c.SetRequest(c.Request().WithContext(ctx))

			// Next
			return next(c)
		}
	}
}
