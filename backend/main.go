package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tocoteron/kankaku/graph"
	"github.com/tocoteron/kankaku/graph/generated"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Get
func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func main() {
	// Get env vars
	port := getEnv("PORT", "8080")

	// Create server
	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set GraphQL routes
	e.POST("/graphql", echo.WrapHandler(graphqlHandler()))
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
