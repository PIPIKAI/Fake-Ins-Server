package routers

import (
	"log"

	"github.com/PIPIKAI/Ins-gin-vue/server/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerConfig(r *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.Host = viper.GetString("server.host")
	docs.SwaggerInfo.BasePath = viper.GetString("server.basepath")
	prefixPath := viper.GetString("swagger.prefix")
	var authorized *gin.RouterGroup
	if viper.GetBool("env.release") {
		authorized = r.Group(prefixPath, gin.BasicAuth(gin.Accounts{
			viper.GetString("swagger.username"): viper.GetString("swagger.password"),
		}))
	} else {
		authorized = r.Group(prefixPath)
	}
	log.Println("swag running at https://" + docs.SwaggerInfo.Host + prefixPath + "/index.html")

	authorized.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
