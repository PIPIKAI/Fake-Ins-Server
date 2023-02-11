package postservice

import (
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/orm"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateImgUrls(urls []string) []model.ImgUrl {
	res := make([]model.ImgUrl, len(urls))
	DB := common.GetDB()
	for idx, url := range urls {
		Mc := model.ImgUrl{}
		if err := DB.Where("url = ?", url).First(&Mc).Error; err != nil {
			res[idx].Url = url
		} else {
			res[idx] = Mc
		}
	}
	return res
}
func CreateCategorys(names []string) []model.Category {
	res := make([]model.Category, len(names))
	for id, v := range names {
		res[id].Name = v
	}
	return res
}

// @Summary     创建帖子
// @Schemes     http https
// @Description Create
// @Param       Form body CreatePostForm true "CreatePostForm"
// @Param       uid  query int false "UID"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/create [post]
func (c PostService) Create(ctx *gin.Context) {
	uid := ctx.Query("uid")
	userId, _ := strconv.Atoi(uid)
	DB := common.GetDB()
	catch, err := c.CreatePost(ctx)
	if err != nil {
		c.ResErr(ctx, "CreatePost:"+err.Error())
		return
	}
	// 创建帖子
	post := &model.Post{
		UserID:    uint(userId),
		Categorys: c.CreateOrFind(catch.Categorys),
		Place:     catch.Place,
		ImgUrls:   CreateImgUrls(catch.Pictures),
		Explain:   catch.MyComment,
	}
	if err := DB.Preload("User").Preload("Post").Create(&post).Error; err != nil {
		util.Response.Error(ctx, nil, err.Error())
		return
	}
	util.Response.Success(ctx, gin.H{"post": post}, "创建成功")

}

// 获得一个帖子
// @Summary     GetByPostID
// @Schemes     http https
// @Description GetByPostID
// @Param       postid path int true "postid"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/getby/postid/{postid} [post]
func (c PostService) GetByPostID(ctx *gin.Context) {

	postID, _ := strconv.Atoi(ctx.Param("postid"))
	post := Post{ID: uint(postID)}
	DB := common.GetDB()
	if err := DB.Preload(clause.Associations).First(&post).Error; err != nil {
		util.Response.Error(ctx, nil, err.Error())
		return
	}
	util.Response.Success(ctx, gin.H{"post": post}, "查询成功")
}

//
// @Summary     EditByPostID
// @Schemes     http https
// @Description 编辑一个帖子
// @Param       form body EditPostForm true "EditPostForm"
// @Param       postid path int true "postid"
// @Param       uid  query int false "UID"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/edit/{postid} [put]
func (c PostService) EditByPostID(ctx *gin.Context) {
	// 获得绑定数据
	DB := common.GetDB()
	postID, _ := strconv.Atoi(ctx.Param("postid"))
	OriginPost := Post{ID: uint(postID)}
	if err := DB.First(&OriginPost).Error; err != nil {
		c.ResErr(ctx, "帖子不存在")
		return
	}
	catch, err := c.EditPost(ctx)
	if err != nil {
		c.ResErr(ctx, "表单错误")
		return
	}
	// 创建帖子
	NewPost := &model.Post{
		UserID:    OriginPost.UserID,
		ID:        OriginPost.ID,
		Categorys: c.CreateOrFind(catch.Categorys),
		Place:     catch.Place,
		Explain:   catch.MyComment,
	}
	if err := DB.Model(model.Post{ID: OriginPost.ID}).Preload(clause.Associations).Updates(&NewPost).Error; err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, nil, "修改成功")

}

//
// 获得该用户的所有帖子
// @Summary     获得该用户的所有帖子
// @Schemes     http https
// @Description GetByUser
// @Param       uid path int true "uid"
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/getby/uid/{uid} [post]
func (c PostService) GetByUser(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	page, pageSize := util.PagePreSolve(ctx)
	var posts []Post
	DB := common.GetDB()
	user := model.User{ID: uint(uid)}
	DB.Scopes(orm.Paginate(page, pageSize)).Order("created_at desc").Model(&user).Preload(clause.Associations).Association("Posts").Find(&posts)
	util.Response.Success(ctx, gin.H{"page": page, "page_size": pageSize, "data": posts}, "查询成功")
}

// 获得所有用户所有帖子
func (c PostService) GetUsersPosts(ids []uint) ([]Post, error) {
	var posts []Post
	DB := common.GetDB()

	err := DB.Order("created_at desc").Model(model.User{}).Where("ID = ?", ids).Preload("Categorys").Preload("ImgUrls").Association("Posts").Find(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 删除一个帖子
// @Summary     DeleteByPostID
// @Schemes     http https
// @Description 删除一个帖子
// @Param       postid path int true "postid"
// @Param       uid query int false "uid"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/delete/{postid} [delete]
func (c PostService) DeleteByPostID(ctx *gin.Context) {
	// 获得绑定数据
	DB := common.GetDB()
	uid := ctx.Query("uid")
	userId, _ := strconv.Atoi(uid)
	postID, _ := strconv.Atoi(ctx.Param("postid"))

	post := model.Post{ID: uint(postID), UserID: uint(userId)}
	if err := DB.Model(model.Post{}).First(&post).Error; err != nil {
		c.ResErr(ctx, "帖子不存在")
		return
	}
	if err := DB.Select("ImgUrls").Delete(&post).Error; err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, nil, "删除成功")
}

// 获得该用户关注的用户所有的帖子
// @Summary     获得该用户关注的用户所有的帖子
// @Schemes     http https
// @Description GetWaths
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/get/home [post]
func (PostService) GetWaths(ctx *gin.Context) {
	tp, _ := ctx.Get("user")
	self := tp.(model.User)
	var posts []Post
	var watchUsers []model.User
	DB := common.GetDB()
	DB.Model(&self).Select("ID").Association("Watchs").Find(&watchUsers)
	watchUsers = append(watchUsers, self)
	page, page_size := util.PagePreSolve(ctx)
	DB.Scopes(orm.Paginate(page, page_size)).Order("created_at desc").Model(&watchUsers).Preload("Comments").Preload("Categorys").Preload("ImgUrls").Association("Posts").Find(&posts)

	util.Response.Success(ctx, gin.H{"data": posts, "page": page, "page_size": page_size}, "查询成功")
}
