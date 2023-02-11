package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解决跨域问题
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://172.27.224.1:3002")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,XFILENAME,XFILECATEGORY,XFILESIZE")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		}
		ctx.Next()
	}
}
