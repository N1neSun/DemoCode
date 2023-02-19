package app

import (
	"test_gin/app/api"
)

type App struct {
	User *api.User
}

func NewApp(user *api.User) *App {
	return &App{
		User: user,
	}
}
