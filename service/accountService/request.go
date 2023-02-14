package accountservice

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/upload"

	"github.com/gin-gonic/gin"
)

type AttemptForm struct {
	Username     string `json:"username" binding:"required" example:"zzk"`
	Name         string `json:"name" binding:"required" example:"志明"`
	Password     string `json:"password" binding:"required" example:"123456"`
	EmailOrPhone string `json:"emailorphone" binding:"required" example:"1652091948@qq.com"`
}
type RegisterForm struct {
	Username     string `json:"username" binding:"required" example:"zzk"`
	Name         string `json:"name" binding:"required" example:"志明"`
	Password     string `json:"password" binding:"required" example:"123456"`
	EmailOrPhone string `json:"emailorphone" binding:"required" example:"1652091948@qq.com"`
	Telephone    string `json:"telephone" swaggerignore:"true" example:""`
	Email        string `json:"email" swaggerignore:"true" example:""`
	BirthDay     string `json:"birth_day" example:"2022-12-12"`
	Code         string `json:"code" example:"1234"`
}

type LoginForm struct {
	Password               string `json:"password" binding:"required" example:"123456"`
	EmailOrPhoneOrUsername string `json:"emailorphoneorusername" binding:"required" example:"zzk"`
}

type PhotoFrom struct {
	ID       uint              `json:"uid"`
	DataList []upload.DataList `json:"data_list"`
}

type EditForm struct {
	ID       uint   `json:"ID" binding:"required"  example:"1"`
	Name     string `json:"Name"  example:"志明"`
	Website  string `json:"Website"  example:"https://example.com"`
	Bio      string `json:"Bio"  example:"my bio"`
	Gender   string `json:"Gender"  example:"mele"`
	BirthDay string `json:"BirthDay"  example:"2022-12-12"`
}

func validatPE(s string) (phone string, email string) {
	phoneReg := regexp.MustCompile(phoneRegex)
	emailReg := regexp.MustCompile(emailRegex)
	isEmail := emailReg.MatchString(s)
	isPhone := phoneReg.MatchString(s)
	if isEmail {
		return "", s
	} else if isPhone {
		return s, ""
	} else {
		return "", ""
	}
}
func (c AccountService) ValidataRegister(ctx *gin.Context) (m *RegisterForm, response gin.H, err error) {
	// var v *model.User
	var catch *RegisterForm

	if err := ctx.ShouldBind(&catch); err != nil {
		log.Println("catch", catch)
		return nil, nil, fmt.Errorf("数据不规范")
	}
	// 验证是手机还是邮箱
	phone, email := validatPE(catch.EmailOrPhone)
	if phone == email {
		return nil, gin.H{"position": 1}, fmt.Errorf("数据不规范")
	}
	catch.Email = email
	catch.Telephone = phone

	DB := common.GetDB()
	v := &model.User{}
	log.Println(phone, email, catch)
	if email != "" {
		if err := DB.Where("email = ?", email).First(&v).Error; err == nil {
			return nil, gin.H{"position": 1}, fmt.Errorf("已存在该邮箱")
		}
	} else {
		if err := DB.Where("telephone = ?", phone).First(&v).Error; err == nil {
			return nil, gin.H{"position": 1}, fmt.Errorf("已存在该手机号")
		}
	}

	if err := DB.Where("username = ?", catch.Username).First(&v).Error; err == nil {
		return nil, gin.H{"position": 2}, fmt.Errorf("已存在该用户名")
	}

	return catch, nil, nil
}
func (c AccountService) ValidataLogin(ctx *gin.Context) (p string, m *model.User, response gin.H, err error) {

	var catch *LoginForm
	if err := ctx.ShouldBind(&catch); err != nil {
		log.Println("catch", catch)
		return "", nil, nil, fmt.Errorf("数据不规范")
	}
	// 验证是手机还是邮箱
	// phone, email := validatPE(catch.EmailOrPhoneOrUsername)
	// if phone == email {
	// 	return "", nil, gin.H{"position": 1}, fmt.Errorf("用户不存在")
	// }
	// catch.Email = email
	// catch.Telephone = phone

	v := &model.User{}
	e := catch.EmailOrPhoneOrUsername
	if err := common.GetDB().Where("email = ? Or user_name = ? Or telephone = ?", e, e, e).First(&v).Error; err != nil {
		return "", nil, nil, fmt.Errorf("用户不存在")
	}
	return catch.Password, v, nil, nil
}
func (c AccountService) ValidataEdit(ctx *gin.Context) (e *EditForm, err error) {

	var catch *EditForm
	if err := ctx.ShouldBind(&catch); err != nil {
		log.Println("catch", catch)
		return nil, fmt.Errorf("数据不规范")
	}
	return catch, nil
}
