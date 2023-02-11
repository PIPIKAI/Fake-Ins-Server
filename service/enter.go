package service

import (
	accountservice "github.com/PIPIKAI/Ins-gin-vue/server/service/accountService"
	commentservice "github.com/PIPIKAI/Ins-gin-vue/server/service/commentService"
	likeservice "github.com/PIPIKAI/Ins-gin-vue/server/service/likeService"
	postservice "github.com/PIPIKAI/Ins-gin-vue/server/service/postService"
	universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"
	userservice "github.com/PIPIKAI/Ins-gin-vue/server/service/userService"
	"github.com/gin-gonic/gin"
)

type Service struct {
	universalservice.UniversalService
	UserService    userservice.UserService
	AccountService accountservice.AccountService
	PostService    postservice.PostService
	LikeService    likeservice.LikeService
	CommentService commentservice.CommentService
	// CategoryService
}

func Register() Service {
	return Service{
		UserService:    userservice.Register(),
		AccountService: accountservice.Register(),
		PostService:    postservice.Register(),
		LikeService:    likeservice.Register(),
		CommentService: commentservice.Register(),
	}
}

func (s Service) Ehco(ctx *gin.Context) {
	s.ResSuccess(ctx, nil, "ehco")
}
