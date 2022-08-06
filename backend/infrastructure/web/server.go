package web

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tocoteron/kankaku/infrastructure/web/auth"
	"github.com/tocoteron/kankaku/interface/app"
	"github.com/tocoteron/kankaku/interface/handler/graphql/generated"
	"github.com/tocoteron/kankaku/interface/handler/graphql/model"
	"github.com/tocoteron/kankaku/interface/handler/graphql/resolver"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	config serverConfig
}

type serverConfig struct {
	app    *app.App
	port   string
	secret []byte
}

func NewServer(app *app.App, port string, secret []byte) *server {
	return &server{
		config: serverConfig{
			app:    app,
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
	e.POST("/signup", signup(s.config.app, s.config.secret))

	// Set GraphQL routes
	e.POST(
		"/graphql",
		graphqlHandler(s.config.app),
		auth.TokenValidator(s.config.secret),
		auth.UserContextProvider(),
	)
	e.GET("/playground", graphqlPlaygroundHandler("/graphql"))

	// Start server
	serverPort := fmt.Sprintf(":%s", s.config.port)
	e.Logger.Fatal(e.Start(serverPort))
}

// GraphQL
func graphqlHandler(app *app.App) echo.HandlerFunc {
	resolver := resolver.NewResolver(app)
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)
	return echo.WrapHandler(h)
}

// GraphQL playground
func graphqlPlaygroundHandler(graphqlPath string) echo.HandlerFunc {
	h := playground.Handler("GraphQL playground", graphqlPath)
	return echo.WrapHandler(h)
}

func signup(app *app.App, secret []byte) echo.HandlerFunc {
	type Req struct {
		Name string `json:"name"`
	}

	type Res struct {
		User  *model.User `json:"user"`
		Token string      `json:"token"`
	}

	return func(c echo.Context) error {
		req := &Req{}

		if err := c.Bind(req); err != nil {
			return err
		}

		u, err := app.UserUseCase().CreateUser(req.Name)
		if err != nil {
			return err
		}

		token, err := auth.GenerateToken(u.ID().String(), secret)
		if err != nil {
			return err
		}

		user := &model.User{
			ID:    u.ID().String(),
			Name:  u.Name(),
			Posts: []*model.Post{},
		}

		return c.JSON(http.StatusOK, Res{
			User:  user,
			Token: token,
		})
	}
}
