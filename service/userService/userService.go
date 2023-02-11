package userservice

import (
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"

	"github.com/gin-gonic/gin"
)

// func (c UserService) ToSelfInfo(DB *gorm.DB) (tx *gorm.DB) {
// 	return DB.Omit(dto.ToSelfOmit...).Preload("Watchs", func(db *gorm.DB) *gorm.DB {
// 		return db.Omit(dto.ToOthersOmit...)
// 	}).Preload("Fans", func(db *gorm.DB) *gorm.DB {
// 		return db.Omit(dto.ToOthersOmit...)
// 	}).Preload("Posts", func(db *gorm.DB) *gorm.DB {
// 		return db.Preload(clause.Associations)
// 	}).Preload("Collections")
// }
// func (c UserService) ToOtherInfo(DB *gorm.DB) (tx *gorm.DB) {
// 	return DB.Omit(dto.ToOthersOmit...).Preload("Watchs", func(db *gorm.DB) *gorm.DB {
// 		return db.Omit(dto.ToOthersOmit...)
// 	}).Preload("Fans", func(db *gorm.DB) *gorm.DB {
// 		return db.Omit(dto.ToOthersOmit...)
// 	}).Preload("Posts", func(db *gorm.DB) *gorm.DB {
// 		return db.Preload(clause.Associations)
// 	}).Preload("Collections")
// }

// @Summary     GetByUserName
// @Schemes     http https
// @Description 获取用户信息
// @Param       username path string true "username"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/getby/username/{username} [post]
func (c UserService) GetByUserName(ctx *gin.Context) {
	username := ctx.Param("username")
	var user RESUser

	if err := common.GetDB().Model(model.User{}).Where("user_name = ?", username).First(&user).Error; err != nil {
		c.ResErr(ctx, "不存在该用户")
		return
	}
	c.ResSuccess(ctx, gin.H{"data": user}, "获取成功")
}

// @Summary     GetByID
// @Schemes     http https
// @Description 获取用户信息
// @Param       uid path int true "uid"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/getby/uid/{uid} [post]
func (c UserService) GetByID(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	var user RESUser
	if err := common.GetDB().Model(model.User{}).Where("id = ?", uid).First(&user).Error; err != nil {
		c.ResErr(ctx, "不存在该用户")
		return
	}
	user.IsWatched = c.IsWatched(c.GetSelfID(ctx), uint(uid))
	c.ResSuccess(ctx, gin.H{"data": user}, "获取成功")
}
func (c UserService) GetByIDS(ids []uint) []RESUser {
	var user []RESUser
	if err := common.GetDB().Model(model.User{}).Where("id = ?", ids).Find(&user).Error; err != nil {
		panic(err.Error())
	}
	return user
}

func (c UserService) List(ctx *gin.Context) {
	uers := &model.User{}
	if err := common.GetDB().Find(&uers).Error; err != nil {
		util.Response.Error(ctx, nil, err.Error())
	}
	util.Response.Success(ctx, gin.H{"data": uers}, "success")
}
