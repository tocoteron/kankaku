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
	e.POST("/graphql", graphqlHandler(), auth.TokenValidator(secret), auth.UserContextProvider())
	e.GET("/playground", graphqlPlaygroundHandler("/graphql"))

	// Start server
	serverPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverPort))
}

// GraphQL
func graphqlHandler() echo.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
	return echo.WrapHandler(h)
}

// GraphQL playground
func graphqlPlaygroundHandler(graphqlPath string) echo.HandlerFunc {
	h := playground.Handler("GraphQL playground", graphqlPath)
	return echo.WrapHandler(h)
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

		token, err := auth.GenerateToken(0, secret)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, Res{Token: token})
	}
}
