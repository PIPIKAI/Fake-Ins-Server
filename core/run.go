package core

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/core/routers"
	"github.com/PIPIKAI/Ins-gin-vue/server/initialize"
)

func Run() {
	initialize.InitAll()
	routers.Run()
}
