package accountservice

import (
	universalservice "github.com/PIPIKAI/Ins-gin-vue/server/service/universalService"
)

type AccountService struct {
	universalservice.UniversalService
}

const emailRegex = `\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`
const phoneRegex = `^1([38]\d|5[0-35-9]|7[3678])\d{8}$`

func Register() AccountService {
	return AccountService{}
}
