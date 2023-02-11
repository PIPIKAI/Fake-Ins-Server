package userservice

import (
	universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"
	"github.com/gin-gonic/gin"
)

type WatchService struct {
	universalservice.UniversalService
}
type IUserService interface {
	GetByID(ctx *gin.Context)
	GetByIDS(Ids []uint) []RESUser
	GetByUserName(ctx *gin.Context)
}

type UserService struct {
	universalservice.UniversalService
	WatchService
}

func Register() UserService {
	return UserService{
		WatchService: WatchService{},
	}
}
