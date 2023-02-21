package routers

import (
	"time"

	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	cache "github.com/chenyahui/gin-cache"

	"github.com/gin-gonic/gin"
)

func PostGroup(group *gin.RouterGroup) *gin.RouterGroup {
	postService := Service.PostService
	likeService := Service.LikeService
	r := group.Group("/post")
	r.Use(middleware.LoginedMiddleware())
	r.POST("/create", middleware.AuthMiddleware(), postService.Create)
	r.POST("/get/home", postService.GetWaths)
	r.POST("/getby/postid/:postid", postService.GetByPostID)
	r.POST("/getby/uid/:uid", postService.GetByUser)
	r.DELETE("/delete/:postid", middleware.AuthMiddleware(), postService.DeleteByPostID)
	r.PUT("/edit/:postid", middleware.AuthMiddleware(), postService.EditByPostID)
	r.POST("/like", likeService.Like)
	r.DELETE("/undolike", likeService.UndoLike)
	r.POST("/likedornot", cache.CacheByRequestURI(CacheStore, 2*time.Second), likeService.LikedOrNot)
	// postRouter.GET("/:id", postController.Show)
	// postRouter.DELETE("/:id", postController.Delete)
	// postRouter.POST("/page/list", postController.PageList)
	return r
}
