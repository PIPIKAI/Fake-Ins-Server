package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FrequentMiddleware() gin.HandlerFunc {
	// 用来防止对某个服务过于频繁的访问，设置时限
	// 例如发邮件
	return func(ctx *gin.Context) {

		now := time.Now().Unix()
		if endtime, exist := ctx.Get("endAt"); exist {
			if endtime.(int64) > now {
				ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": fmt.Sprintf("操作过于频繁，请在 %vs后重试", endtime.(int64)-now)})
				// 直接放弃next
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
