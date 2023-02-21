package routers

import (
	"time"

	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

func UserGroup(group *gin.RouterGroup) *gin.RouterGroup {
	userService := Service.UserService
	watchService := userService.WatchService
	r := group.Group("/user")
	r.Use(middleware.LoginedMiddleware())

	r.POST("/info", cache.CacheByRequestURI(CacheStore, 1*time.Minute), Service.AccountService.Info)
	r.PUT("/info", Service.AccountService.EditCountInfo)
	r.POST("/logout", Service.AccountService.Logout)
	r.POST("/photo", Service.AccountService.ChangePhoto)
	// r.POST("/commend/users", userService.CommentsUsers)
	r.POST("/getby/username/:username", cache.CacheByRequestURI(CacheStore, 2*time.Second), userService.GetByUserName)
	r.POST("/getby/uid/:uid", cache.CacheByRequestURI(CacheStore, 2*time.Second), userService.GetByID)
	r.POST("/watch/:uid", watchService.WatchUser)
	r.POST("/unwatch/:uid", watchService.UnWatchUser)
	r.POST("/unwatchedusers", watchService.UnWatchedUsers)
	r.POST("/watchedusers", watchService.HadWatchedUsers)
	r.POST("/watchedornot", watchService.WatchedOrNot)
	r.POST("/bewatchedornot", watchService.BeWatchedOrNot)
	r.POST("/getfans", watchService.GetFans)

	return r
}
