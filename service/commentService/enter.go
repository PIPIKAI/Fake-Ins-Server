package commentservice

import universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"

type CommentService struct {
	universalservice.UniversalService
}

func Register() CommentService {
	return CommentService{}
}
