package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func TimerRecordMiddleware() gin.HandlerFunc {
	// 用来防止对某个服务过于频繁的访问，设置时限
	// 例如发邮件
	return func(ctx *gin.Context) {
		now := time.Now().Unix()
		ctx.Set("beginAt", now)
		ctx.Next()
	}
}
