package routers

import (
	"github.com/gin-gonic/gin"
)

func __(r *gin.Engine) *gin.Engine {
	// r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	// apiRouter := r.Group("/api")

	// apiRouter.Use(sessions.SessionsMany([]string{"info", "mid"}, common.GetRedis()))

	// userController := controller.NewUserController()

	// apiRouter.POST("/register/attempt", userController.Attempt)
	// apiRouter.POST("/register/sendmailcode", userController.PostEmail)
	// apiRouter.POST("/register/register", userController.Register)
	// apiRouter.POST("/user/login", userController.Login)

	// // apiRouter.POST("/login", userController.Login)

	// authApiRouter := apiRouter.Group("/auth")
	// authApiRouter.Use(middleware.LoginedMiddleware())

	// userRouter := authApiRouter.Group(("/user"))

	// userRouter.POST("/info", userController.Info)
	// userRouter.POST("/logout", userController.Logout)
	// userRouter.POST("/commend/users", userController.CommentsUsers)
	// userRouter.POST("/getby/username/:username", userController.GetUsersByUserName)
	// userRouter.POST("/watch/:uid", userController.WatchUser)
	// userRouter.POST("/unwatch/:uid", userController.UnWatchUser)

	// postRouter := authApiRouter.Group("/post")
	// postController := controller.NewPostController()
	// postRouter.POST("/create", postController.Create)
	// postRouter.POST("/get/home", postController.GetWaths)
	// postRouter.POST("/getby/postid/:postid", postController.GetByPostID)
	// postRouter.POST("/getby/uid/:uid", postController.GetByUser)
	// postRouter.DELETE("/delete/:postid", postController.DeleteByPostID)
	// // postRouter.PUT("/:id", postController.Update)
	// // postRouter.GET("/:id", postController.Show)
	// // postRouter.DELETE("/:id", postController.Delete)
	// // postRouter.POST("/page/list", postController.PageList)

	return r
}
