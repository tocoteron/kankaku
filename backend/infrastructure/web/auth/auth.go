package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tocoteron/kankaku/infrastructure/web/graphql"
)

const tokenContextKey = "token"

type JWTCustomClaims struct {
	jwt.StandardClaims
}

// Authenticate user by id and password
func Authenticate(id string, password string) bool {
	fmt.Println(id, password)
	if id == "test" && password == "password" {
		return true
	}
	return false
}

// Generate token to authenticate user
func GenerateToken(id uint64, secret []byte) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		jwt.StandardClaims{
			Id:        strconv.FormatUint(id, 10),
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
				return fmt.Errorf("failed to retrieve token from context")
			}
			token, ok := jwtCtx.(*jwt.Token)
			if !ok {
				return fmt.Errorf("failed to convert token to *jwt.Token")
			}
			claims, ok := token.Claims.(*JWTCustomClaims)
			if !ok {
				return fmt.Errorf("failed to convert jwt claims to *auth.JWTCustomClaims")
			}

			// Parse user id of token
			id, err := strconv.ParseUint(claims.Id, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse token id: %w", err)
			}

			// Set user context
			ctx := graphql.SetUserContext(c.Request().Context(), &graphql.UserContext{ID: id})
			c.SetRequest(c.Request().WithContext(ctx))

			// Next
			return next(c)
		}
	}
}
