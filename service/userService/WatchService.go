package userservice

import (
	"log"
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/orm"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c WatchService) IsWatched(selfID, uid uint) bool {
	DB := common.GetDB()
	var watchs []model.User
	DB.Model(&model.User{ID: selfID}).Where("ID = ?", uid).Association("Watchs").Find(&watchs)
	return len(watchs) != 0
}

func (c WatchService) IsBeWatched(selfID, uid uint) bool {
	DB := common.GetDB()
	var watchs []model.User
	DB.Model(&model.User{ID: uid}).Where("ID = ?", selfID).Association("Watchs").Find(&watchs)
	return len(watchs) != 0
}

// @Summary     是否关注
// @Schemes     http https
// @Description WatchedOrNot
// @Param       uid query int true "uid"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/watchedornot [post]
func (c WatchService) WatchedOrNot(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Query("uid"))
	if c.IsWatched(c.GetSelfID(ctx), uint(uid)) {
		c.ResSuccess(ctx, gin.H{"data": true}, "已经关注")
	} else {
		c.ResSuccess(ctx, gin.H{"data": false}, "未关注")
	}
}

// @Summary     是否关注
// @Schemes     http https
// @Description BeWatchedOrNot
// @Param       uid query int true "uid"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/bewatchedornot [post]
func (c WatchService) BeWatchedOrNot(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Query("uid"))
	if c.IsBeWatched(c.GetSelfID(ctx), uint(uid)) {
		c.ResSuccess(ctx, gin.H{"data": true}, "被关注")
	} else {
		c.ResSuccess(ctx, gin.H{"data": false}, "未被关注")
	}
}

// @Summary     WatchUser
// @Schemes     http https
// @Description 关注用户
// @Param       uid path int true "uid"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/watch/{uid} [post]
func (c WatchService) WatchUser(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	selfID := c.GetSelfID(ctx)

	if selfID == uint(uid) {
		c.ResErr(ctx, "不能自己关注自己")
		return
	}
	DB := common.GetDB()
	var user1, user2 model.User
	var watchs []model.User
	if err1, err2 := DB.First(&user1, selfID).Error, DB.First(&user2, uint(uid)).Error; err1 != nil || err2 != nil {
		c.ResErr(ctx, "用户不存在")
		return
	}

	// 如果已经关注了
	//
	DB.Model(&user1).Where("ID = ?", user2.ID).Association("Watchs").Find(&watchs)
	if c.IsWatched(selfID, uint(uid)) {
		c.ResErr(ctx, "已经关注了")
		return
	}
	DB.Model(&user1).Association("Watchs").Append(&user2)
	DB.Model(&user1).UpdateColumn("watchs_counts", gorm.Expr("watchs_counts + ?", 1))
	DB.Model(&user2).UpdateColumn("fans_counts", gorm.Expr("fans_counts + ?", 1))

	util.Response.Success(ctx, nil, "关注成功")

}

// @Summary     UnWatchUser
// @Schemes     http https
// @Description 取关用户
// @Param       uid path int true "uid"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/unwatch/{uid} [post]
func (c WatchService) UnWatchUser(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	selfID := c.GetSelfID(ctx)

	if selfID == uint(uid) {
		c.ResErr(ctx, "系统错误")
		return
	}
	DB := common.GetDB()
	var user1, user2 model.User
	if err1, err2 := DB.First(&user1, selfID).Error, DB.First(&user2, uint(uid)).Error; err1 != nil || err2 != nil {
		c.ResErr(ctx, "用户不存在")
		return
	}
	if !c.IsWatched(selfID, uint(uid)) {
		c.ResErr(ctx, "未关注，不能取关")
		return
	}
	if err := DB.Model(&user1).Association("Watchs").Delete(&user2); err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	DB.Model(&user1).UpdateColumn("watchs_counts", gorm.Expr("watchs_counts - ?", 1))

	// DB.Model(&user2).Association("Fans").Delete(&user1)
	DB.Model(&user2).UpdateColumn("fans_counts", gorm.Expr("fans_counts - ?", 1))

	util.Response.Success(ctx, nil, "取关成功")
}

// @Summary     UnWatchedUsers
// @Schemes     http https
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Description 没关注的用户
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/unwatchedusers [post]
func (c WatchService) UnWatchedUsers(ctx *gin.Context) {
	DB := common.GetDB()
	page, pageSize := util.PagePreSolve(ctx)
	selfID := c.GetSelfID(ctx)
	users := []RESUser{}
	watchedUsers := c.WatchedUsers(ctx)
	watchedUsers = append(watchedUsers, selfID)
	log.Println("watchedUsers: ", watchedUsers)
	if err := DB.Scopes(orm.Paginate(page, pageSize)).Order("created_at desc").Model(model.User{}).Not(watchedUsers).Find(&users).Error; err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, gin.H{"page": page, "page_size": pageSize, "data": users}, "success")
}

// @Summary     已经关注的用户
// @Schemes     http https
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Description HadWatchedUsers
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/watchedusers [post]
func (c WatchService) HadWatchedUsers(ctx *gin.Context) {
	DB := common.GetDB()
	page, pageSize := util.PagePreSolve(ctx)
	tp, _ := ctx.Get("user")
	self := tp.(model.User)
	var watchUsers []RESUser
	DB.Scopes(orm.Paginate(page, pageSize)).Order("created_at desc").Model(&self).Association("Watchs").Find(&watchUsers)
	c.ResSuccess(ctx, gin.H{"page": page, "page_size": pageSize, "data": watchUsers}, "success")
}
func (c WatchService) WatchedUsers(ctx *gin.Context) []uint {
	DB := common.GetDB()
	tp, _ := ctx.Get("user")
	self := tp.(model.User)
	var watchUsers []uint
	DB.Model(&self).Select("id").Association("Watchs").Find(&watchUsers)
	return watchUsers
}

// @Summary     获取粉丝列表
// @Schemes     http https
// @Description 获取粉丝列表
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Tags        User
// @Accept      json
// @Produce     json
// @success     200
// @Router      /user/getfans [post]
func (c WatchService) GetFans(ctx *gin.Context) {
	DB := common.GetDB()
	page, pageSize := util.PagePreSolve(ctx)
	tp, _ := ctx.Get("user")
	self := tp.(model.User)
	var fansId []uint
	var fansList []RESUser

	DB.Model(&self).Association("Fans").Find(&fansList)

	DB.Raw("SELECT user_id FROM user_watchs WHERE watch_id = ?", self.ID).Scan(&fansId)

	DB.Scopes(orm.Paginate(page, pageSize)).Order("created_at desc").Model(&model.User{}).Find(&fansList, fansId)
	c.ResSuccess(ctx, gin.H{"page": page, "page_size": pageSize, "data": fansList}, "success")
}

// // 获取未关注的用户
// func (c WatchService) CommentsUsers(ctx *gin.Context) {
// 	self, _ := ctx.Get("user")
// 	var users []RESUser
// 	watchedUsers := c.WatchedUsers(ctx)
// 	if err := common.GetDB().Model(model.User{}).Not(&watchedUsers).Where("ID != ?", self.(model.User).ID).Select([]string{"ID", "UserName", "Name", "Bio", "Website", "Photo"}).Find(&users).Error; err != nil {
// 		util.Response.Error(ctx, nil, err.Error())
// 	}
// 	util.Response.Success(ctx, gin.H{"data": users}, "success")
// }
