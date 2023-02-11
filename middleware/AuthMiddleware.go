package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// 定义auth的中间件 用来检查jwt权限

	return func(ctx *gin.Context) {
		tp, _ := ctx.Get("user")
		self := tp.(model.User)
		// Query("uid") 你要操作的那个uid
		// selfID 自己的uid
		uid, err := strconv.Atoi(ctx.Query("uid"))
		if err != nil || uint(uid) != self.ID {
			ctx.Abort()
			util.Response.ResponsFmt(ctx, http.StatusUnauthorized, 401, nil, "没有权限")
		}
		ctx.Next()
	}
}
func LoginedMiddleware() gin.HandlerFunc {
	// 定义auth的中间件 用来检查jwt权限

	return func(ctx *gin.Context) {
		authsession := sessions.DefaultMany(ctx, "info")
		if info := authsession.Get("user"); info == nil {
			ctx.Abort()
			util.Response.ResponsFmt(ctx, http.StatusUnauthorized, 401, nil, "请先登录")
			return
		} else {
			var userinfo model.User
			if err := json.Unmarshal([]byte(info.(string)), &userinfo); err != nil {
				ctx.Abort()
				util.Response.Error(ctx, nil, err.Error())

				return
			}
			ctx.Set("user", userinfo)
		}
		ctx.Next()
	}
}
