package core

import (
	"log"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/email"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/upload"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitAll() {
	InitViper()

	if viper.GetBool("env.release") {
		log.Println("[warn] Gin Mode : RELEASE")
		gin.SetMode(gin.ReleaseMode)
	}
	common.InitDB()
	common.InitRedis()
	upload.NewQiniuUper()
	email.NewEmailService()
}
