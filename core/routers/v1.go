package routers

import "github.com/gin-gonic/gin"

func V1Group(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/api/v1")

	AccountService := Service.AccountService
	v1.GET("/ehco", Service.Ehco)
	v1.POST("/register/attempt", AccountService.Attempt)
	v1.POST("/register/sendmailcode", AccountService.PostEmail)
	v1.POST("/register/register", AccountService.Register)
	v1.POST("/login", AccountService.Login)
	return v1
}
