package app

import (
	"test_gin/app/api"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewApp)

type App struct {
	User *api.User
}

func NewApp(user *api.User) *App {
	return &App{
		User: user,
	}
}
