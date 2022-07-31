package web

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tocoteron/kankaku/infrastructure/web/auth"
	"github.com/tocoteron/kankaku/infrastructure/web/graph"
	"github.com/tocoteron/kankaku/infrastructure/web/graph/generated"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	config serverConfig
}

type serverConfig struct {
	port   string
	secret []byte
}

func NewServer(port string, secret []byte) *server {
	return &server{
		config: serverConfig{
			port:   port,
			secret: secret,
		},
	}
}

func (s *server) Run() {
	// Create server
	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set auth routes
	e.POST("/signup", signup(s.config.secret))

	// Set GraphQL routes
	e.POST("/graphql", graphqlHandler(), auth.TokenValidator(s.config.secret), auth.UserContextProvider())
	e.GET("/playground", graphqlPlaygroundHandler("/graphql"))

	// Start server
	serverPort := fmt.Sprintf(":%s", s.config.port)
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
