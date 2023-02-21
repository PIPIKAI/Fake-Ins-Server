package routers

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var CacheStore = middleware.GetCacheStore()

func V1Group(r *gin.Engine) {
	v1 := r.Group(viper.GetString("server.basepath"))
	AccountGroup(v1)
	UserGroup(v1)
	PostGroup(v1)
	CommentGroup(v1)
}
