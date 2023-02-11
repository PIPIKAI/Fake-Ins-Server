package middleware

import (
	"fmt"

	"github.com/PIPIKAI/Ins-gin-vue/server/util"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Response.Error(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
