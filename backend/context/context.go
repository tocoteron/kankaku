package context

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
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
