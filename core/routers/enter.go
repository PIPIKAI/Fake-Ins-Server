package routers

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/docs"
	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	"github.com/PIPIKAI/Ins-gin-vue/server/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Service = service.Register()

type Routers struct {
}

func CollectRoute(r *gin.Engine) *gin.Engine {

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := V1Group(r)
	UserGroup(v1)
	PostGroup(v1)
	CommentGroup(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func Run() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.Use(sessions.SessionsMany([]string{"info", "mid"}, common.GetRedis()))

	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(port))
}
