package routers

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	"github.com/PIPIKAI/Ins-gin-vue/server/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Service = service.Register()

func CollectRoute(r *gin.Engine) *gin.Engine {
	V1Group(r)

	return r
}

func Run() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.Use(sessions.SessionsMany([]string{"info", "mid"}, common.GetRedis()))
	r = SwaggerConfig(r)
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(port))
}
