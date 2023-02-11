package postservice

import (
	universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"
)

type CategoryService struct {
}

type PostService struct {
	universalservice.UniversalService
	CategoryService
}

func Register() PostService {
	return PostService{
		CategoryService: CategoryService{},
	}
}
