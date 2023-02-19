//go:build wireinject
// +build wireinject

//go:generate wire
package boot

import (
	"test_gin/app"
	"test_gin/app/api"
	"test_gin/app/repository"
	service "test_gin/app/services"
	"test_gin/common/datasource"
	"test_gin/common/logger"
	"test_gin/common/middleware/jwt"

	"github.com/google/wire"
)

func Inject(source *datasource.Db, log *logger.Logger) (myapp *app.App) {
	panic(wire.Build(
		wire.Struct(new(jwt.JWT), "*"),
		repository.NewRoleRepository,
		repository.NewUserRepository,
		service.NewUserService,
		service.NewRoleService,
		api.NewUser,
		wire.Struct(new(app.App), "*"),
	))
}
