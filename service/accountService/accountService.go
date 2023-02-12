package accountservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	Email "github.com/PIPIKAI/Ins-gin-vue/server/util/email"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary     Register
// @Schemes     http https
// @Description 注册用户
// @Param       Form body RegisterForm true "RegisterForm"
// @Tags        Account
// @Accept      json
// @Produce     json
// @success     200
// @Router      /register/register [post]
func (c AccountService) Register(ctx *gin.Context) {

	// request
	v, respon, err := c.ValidataRegister(ctx)
	if err != nil {
		util.Response.Error(ctx, respon, err.Error())
		return
	}
	// logic

	session := sessions.DefaultMany(ctx, "mid")

	emailInter := session.Get("email")
	if emailInter == nil {
		util.Response.Error(ctx, nil, "emailInter = nil")
		return
	}

	codeInter := session.Get(emailInter.(string))

	if codeInter == nil {
		util.Response.Error(ctx, nil, "codeInter = nil")
		return

	}
	email := emailInter.(string)
	code := codeInter.(string)
	if v.Email != email || v.Code != code {
		util.Response.Error(ctx, nil, "验证码错误")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v.Password), bcrypt.DefaultCost)

	if err != nil {
		util.Response.ResponsFmt(ctx, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}
	log.Println(v)
	user := &model.User{
		UserName:  v.Username,
		Name:      v.Name,
		PassWord:  string(hashedPassword),
		Email:     v.Email,
		BirthDay:  v.BirthDay,
		Telephone: v.Telephone,
	}
	if err := common.GetDB().Create(&user).Error; err != nil {
		util.Response.Error(ctx, nil, err.Error())
		return
	}

	session.Delete("email")
	session.Options(sessions.Options{
		MaxAge:   -1,
		SameSite: 4,
		Secure:   true,
	})
	session.Delete(email)
	session.Save()

	// 发放token
	authsession := sessions.DefaultMany(ctx, "info")
	defer authsession.Save()
	authsession.Options(sessions.Options{
		MaxAge:   3600 * 24 * 7,
		Secure:   true,
		SameSite: 4,
	})
	user.PassWord = ""
	if userinfo, err := json.Marshal(user); err != nil {
		util.Response.Error(ctx, nil, "系统异常")
	} else {
		authsession.Set("user", string(userinfo))
	}

	// response
	util.Response.Success(ctx, gin.H{"data": "", "cookie": ctx.Request.Header.Get("Cookie")}, "注册成功")
}

// @Summary     PostEmail
// @Schemes     http https
// @Description 发送邮件
// @Tags        Account
// @Accept      json
// @Produce     json
// @success     200
// @Router      /register/sendmailcode [post]
func (c AccountService) PostEmail(ctx *gin.Context) {
	session := sessions.DefaultMany(ctx, "mid")
	defer session.Save()

	now := time.Now().Unix()
	if endtime := session.Get("endEmailAt"); endtime != nil {
		if endtime.(int64) > now {
			util.Response.Error(ctx, nil, fmt.Sprintf("操作过于频繁，请在 %vs后重试", endtime.(int64)-now))
			return
		}
	}

	emailInter := session.Get("email")
	phoneInter := session.Get("telephone")
	// telephone := session.Get("telephone").(string)
	// if user.Email != "" {
	if emailInter == nil && phoneInter == nil {
		util.Response.Error(ctx, nil, "邮箱或电话过期，请重试")
		return
	}
	if phoneInter.(string) != "" {
		util.Response.Error(ctx, nil, "暂不支持手机注册")
		return
	}
	email := emailInter.(string)
	ms := Email.GetMailS()
	err, code := ms.SendValidCode(email)
	if err != nil {
		panic(err.Error())
	}
	// 使用session
	// 设置session 的过期时间
	session.Options(sessions.Options{
		MaxAge:   60 * 5,
		Secure:   true,
		SameSite: 4,
	})
	// 设置session
	session.Set(email, code)

	endtime := time.Now().Add(time.Second * 20).Unix()
	session.Set("endEmailAt", endtime)
	log.Println("session Set(endEmailAt)", endtime)
	// 验证邮箱
	// code
	util.Response.Success(ctx, gin.H{"email": email, "code": code}, "发送验证码成功")
}

// @Summary     Attempt
// @Schemes     http https
// @Description 验证注册表单
// @Param       Form body AttemptForm true "AttemptForm"
// @Tags        Account
// @Accept      json
// @Produce     json
// @success     200
// @Router      /register/attempt [post]
func (c AccountService) Attempt(ctx *gin.Context) {
	// 验证数据
	v, respon, err := c.ValidataRegister(ctx)
	if err != nil {
		util.Response.Error(ctx, respon, err.Error())
		return
	}
	// 使用session
	session := sessions.DefaultMany(ctx, "mid")
	// 设置session 的过期时间
	session.Options(sessions.Options{
		MaxAge:   60 * 10,
		Secure:   true,
		SameSite: 4,
	})

	// 设置session
	session.Set("email", v.Email)
	session.Set("telephone", v.Telephone)
	session.Save()

	util.Response.Success(ctx, gin.H{"data": "ok", "v": v}, "ok")
}

// @Summary     Login
// @Schemes     http https
// @Description 登录
// @Param       Form body LoginForm true "LoginForm"
// @Tags        Account
// @Accept      json
// @Produce     json
// @success     200
// @Router      /login [post]
func (c AccountService) Login(ctx *gin.Context) {

	password, user, respon, err := c.ValidataLogin(ctx)
	if err != nil {
		util.Response.Error(ctx, respon, err.Error())
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
		util.Response.Error(ctx, nil, "密码错误")
		return
	}
	// 发放token
	authsession := sessions.DefaultMany(ctx, "info")
	authsession.Options(sessions.Options{
		MaxAge:   3600 * 7 * 24,
		Secure:   true,
		SameSite: 4,
	})

	user.PassWord = ""
	log.Println("user", user)
	if userinfo, err := json.Marshal(user); err != nil {
		util.Response.Error(ctx, nil, "系统异常")
	} else {
		authsession.Set("user", string(userinfo))
	}
	authsession.Save()
	util.Response.Success(ctx, gin.H{"cookie": ctx.Writer.Header().Values("Set-Cookie")[0]}, "登录成功")
}

// @Summary     Logout
// @Schemes     http https
// @Description 登出
// @Tags        Account
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/logout [post]
func (c AccountService) Logout(ctx *gin.Context) {
	session := sessions.DefaultMany(ctx, "info")
	session.Delete("user")
	session.Options(sessions.Options{
		MaxAge:   -1,
		Secure:   true,
		SameSite: 4,
	})
	session.Save()
	c.ResSuccess(ctx, nil, "注销成功")

}

// @Summary     Info
// @Schemes     http https
// @Description 获取用户登录信息
// @Tags        User
// @param info header string false "info"
// @Produce     json
// @success     200
// @Security ApiKeyAuth
// @Router      /user/info [post]
func (c AccountService) Info(ctx *gin.Context) {
	DB := common.GetDB()
	selfUserID := c.GetSelfID(ctx)
	log.Println(selfUserID)
	var user UserDto
	if err := DB.Model(model.User{}).First(&user, selfUserID).Error; err != nil {
		// util.Response.Error(ctx, nil, "获取失败")
		// 清除缓存
		c.Logout(ctx)
		return
	}
	c.ResSuccess(ctx, gin.H{"data": user}, "")

}

// @Summary     Info
// @Schemes     http https
// @Description 获取用户登录信息
// @Tags        User
// @param form body EditForm true "EditForm"
// @Produce     json
// @success     200
// @Router      /user/info [put]
func (c AccountService) EditCountInfo(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	e, err := c.ValidataEdit(ctx)
	if err != nil {
		c.ResErr(ctx, "表单错误")
	}
	if err := common.GetDB().Updates(
		&model.User{
			ID:       uint(uid),
			Bio:      e.Bio,
			BirthDay: e.BirthDay,
			Photo:    e.Photourl,
			Website:  e.Website,
			Gender:   e.Gender,
			Name:     e.Name,
			UserName: e.UserName,
		},
	).Error; err != nil {
		util.Response.Error(ctx, nil, "更新失败")
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, nil, "更新成功")

}
