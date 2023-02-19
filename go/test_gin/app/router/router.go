package router

import (
	"log"
	"test_gin/app/models"
	"test_gin/boot"
	"test_gin/common/datasource"
	"test_gin/common/logger"
	"test_gin/common/middleware/jwt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	// gin.SetMode(setting.Config.APP.RunMode)
	Configure(r)
	return r
}

func Configure(r *gin.Engine) {
	db := datasource.Db{}
	zap := logger.Logger{}
	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	init_tables(db)
	app := boot.Inject(&db, &zap)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var authMiddleware = app.User.Jwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/info", app.User.GetUserInfo)
		userAPI.POST("/logout", app.User.Logout)
		userAPI.GET("/user/list", app.User.GetUsers)
		userAPI.POST("/user", app.User.AddUser)
	}
}

func init_tables(db datasource.Db) error {
	M := db.Conn.Migrator()
	M.AutoMigrate(&models.User{})
	M.AutoMigrate(&models.Role{})
	return nil
}
