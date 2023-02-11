package routers

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/middleware"
	"github.com/gin-gonic/gin"
)

func CommentGroup(group *gin.RouterGroup) *gin.RouterGroup {
	commentService := Service.CommentService
	likeService := Service.LikeService
	r := group.Group("/comment")
	r.Use(middleware.LoginedMiddleware())
	r.GET("/get", commentService.GetComments)
	r.POST("/post", middleware.AuthMiddleware(), commentService.CommentPost)
	r.POST("/reply", middleware.AuthMiddleware(), commentService.ReplyComment)
	r.POST("/like", middleware.AuthMiddleware(), likeService.Like)

	return r
}
