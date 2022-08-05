package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mycontext "github.com/tocoteron/kankaku/interface/handler/context"
)

const tokenContextKey = "token"

type JWTCustomClaims struct {
	jwt.StandardClaims
}

// Authenticate user by id and password
func Authenticate(id string, password string) bool {
	if id == "test" && password == "password" {
		return true
	}
	return false
}

// Generate token to authenticate user
func GenerateToken(id string, secret []byte) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Encode token
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

// Validate token
func TokenValidator(secret []byte) echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		ContextKey: tokenContextKey,
		Claims:     &JWTCustomClaims{},
		SigningKey: secret,
	}
	return middleware.JWTWithConfig(jwtConfig)
}

// Bind user context to context.Context of http.Request
func UserContextProvider() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Retrieve token from context
			jwtCtx := c.Get(tokenContextKey)
			if jwtCtx == nil {
				return fmt.Errorf("failed to get token from context")
			}
			token, ok := jwtCtx.(*jwt.Token)
			if !ok {
				return fmt.Errorf("failed to convert token to *jwt.Token")
			}
			claims, ok := token.Claims.(*JWTCustomClaims)
			if !ok {
				return fmt.Errorf("failed to convert jwt claims to *auth.JWTCustomClaims")
			}

			// Set user context
			ctx := mycontext.SetUserContext(c.Request().Context(), &mycontext.UserContext{ID: claims.Id})
			c.SetRequest(c.Request().WithContext(ctx))

			// Next
			return next(c)
		}
	}
}
