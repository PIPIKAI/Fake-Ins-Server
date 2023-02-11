package likeservice

import universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"

type LikeService struct {
	universalservice.UniversalService
}

func Register() LikeService {
	return LikeService{}
}
