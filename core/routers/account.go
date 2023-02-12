package routers

import (
	"github.com/gin-gonic/gin"
)

func AccountGroup(group *gin.RouterGroup) {
	AccountService := Service.AccountService
	r := group.Group("/")
	r.GET("/ehco", Service.Ehco)
	r.POST("/register/attempt", AccountService.Attempt)
	r.POST("/register/sendmailcode", AccountService.PostEmail)
	r.POST("/register/register", AccountService.Register)
	r.POST("/login", AccountService.Login)
}
