package app

import (
	"github.com/tocoteron/kankaku/domain/service"
	"github.com/tocoteron/kankaku/interface/repository"
	"github.com/tocoteron/kankaku/usecase"
)

type App struct {
	userUseCase usecase.UserUseCase
}

func NewTestApp() *App {
	userRepo := repository.NewUserInMemoryRepository()
	userService := service.NewUserService(userRepo)
	userUsecase := usecase.NewUserUseCase(userService, userRepo)

	return &App{
		userUseCase: userUsecase,
	}
}

func (a *App) UserUseCase() usecase.UserUseCase {
	return a.userUseCase
}
