package core

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/core/routers"
)

func Run() {
	InitAll()
	routers.Run()
}
