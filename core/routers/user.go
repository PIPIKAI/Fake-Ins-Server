package routers

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	"github.com/gin-gonic/gin"
)

func UserGroup(group *gin.RouterGroup) *gin.RouterGroup {
	userService := Service.UserService
	watchService := userService.WatchService
	r := group.Group("/user")
	r.Use(middleware.LoginedMiddleware())

	r.POST("/info", Service.AccountService.Info)
	r.POST("/logout", Service.AccountService.Logout)
	// r.POST("/commend/users", userService.CommentsUsers)
	r.POST("/getby/username/:username", userService.GetByUserName)
	r.POST("/getby/uid/:uid", userService.GetByID)
	r.POST("/watch/:uid", watchService.WatchUser)
	r.POST("/unwatch/:uid", watchService.UnWatchUser)
	r.POST("/unwatchedusers", watchService.UnWatchedUsers)
	r.POST("/watchedusers", watchService.HadWatchedUsers)
	r.POST("/watchedornot", watchService.WatchedOrNot)
	r.POST("/bewatchedornot", watchService.BeWatchedOrNot)
	r.POST("/getfans", watchService.GetFans)

	return r
}
