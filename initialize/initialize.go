package initialize

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/email"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/upload"
)

func InitAll() {
	InitViper()
	common.InitDB()
	common.InitRedis()
	upload.NewQiniuUper()
	email.NewEmailService()
}
