package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tocoteron/kankaku/auth"
	"github.com/tocoteron/kankaku/graph"
	"github.com/tocoteron/kankaku/graph/generated"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Get environment var or default value
func getEnvOrDefault(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// Get environment var or panic program
func getEnvOrPanic(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	panic(fmt.Sprintf("%s is not set", key))
}

func main() {
	// Get env vars
	port := getEnvOrDefault("PORT", "8080")
	secret := []byte(getEnvOrPanic("SECRET"))

	// Create server
	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set auth routes
	e.POST("/signup", signup(secret))

	// Set GraphQL routes
	jwtConfig := middleware.JWTConfig{
		Claims:     &auth.JWTCustomClaims{},
		SigningKey: secret,
	}
	e.POST("/graphql", echo.WrapHandler(graphqlHandler()), middleware.JWTWithConfig(jwtConfig))
	e.GET("/playground", echo.WrapHandler(graphqlPlaygroundHandler("/graphql")))

	// Start server
	serverPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverPort))
}

// GraphQL
func graphqlHandler() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
}

// GraphQL playground
func graphqlPlaygroundHandler(graphqlPath string) http.HandlerFunc {
	return playground.Handler("GraphQL playground", graphqlPath)
}

func signup(secret []byte) echo.HandlerFunc {
	type Req struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	type Res struct {
		Token string `json:"token"`
	}

	return func(c echo.Context) error {
		req := &Req{}

		if err := c.Bind(req); err != nil {
			return err
		}

		if !auth.Authenticate(req.ID, req.Password) {
			return echo.ErrUnauthorized
		}

		token, err := auth.GenerateToken(secret)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, Res{Token: token})
	}
}
