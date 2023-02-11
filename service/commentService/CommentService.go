package commentservice

import (
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/orm"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-gonic/gin"
)

// @Summary     获取评论
// @Schemes     http https
// @Description 获取评论
// @Param       postid  query int false "PID"
// @Param       page query int false "page"
// @Param       page_size query int false "page_size"
// @Tags        Comment
// @Accept      json
// @Produce     json
// @success     200
// @Router      /comment/get [get]
func (c CommentService) GetComments(ctx *gin.Context) {
	postid, _ := strconv.Atoi(ctx.Query("postid"))
	page, pageSize := util.PagePreSolve(ctx)
	post := &model.Post{ID: uint(postid)}
	var comments []Comment
	DB := common.GetDB()
	DB.Scopes(orm.Paginate(page, pageSize)).Order("created_at desc").Model(&post).Where("reply_id is NULL").Preload("Replys").Association("Comments").Find(&comments)
	c.ResSuccess(ctx, gin.H{"page": page, "page_size": pageSize, "data": comments}, "获取评论")

}

// @Summary     评论帖子
// @Schemes     http https
// @Description 评论帖子
// @Param       Form body Comment true "Comment"
// @Param       uid  query int false "UID"
// @Tags        Comment
// @Accept      json
// @Produce     json
// @success     200
// @Router      /comment/post [post]
func (c CommentService) CommentPost(ctx *gin.Context) {
	DB := common.GetDB()
	var catch *ReqComment
	if err := ctx.ShouldBind(&catch); err != nil {
		c.ResErr(ctx, "数据验证错误")
	}
	// 创建评论
	comment := &model.Comment{
		UserID:  uint(catch.UserID),
		PostID:  uint(catch.PostID),
		Content: catch.Content,
	}
	if err := DB.Preload("User").Preload("Post").Preload("Comment").Create(&comment).Error; err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, gin.H{"data": comment}, "评论成功！")

}

// @Summary     评论回复
// @Schemes     http https
// @Description 评论回复
// @Param       Form body Comment true "Comment"
// @Param       uid  query int false "UID"
// @Tags        Comment
// @Accept      json
// @Produce     json
// @success     200
// @Router      /comment/reply [post]
func (c CommentService) ReplyComment(ctx *gin.Context) {
	DB := common.GetDB()
	var catch *ReqReply
	if err := ctx.ShouldBind(&catch); err != nil {
		c.ResErr(ctx, "数据验证错误")
	}
	// 创建回复

	comment := &model.Comment{
		UserID:  uint(catch.UserID),
		PostID:  uint(catch.PostID),
		Content: catch.Content,
		ReplyID: &catch.ReplyID,
	}
	if err := DB.Preload("User").Preload("Post").Preload("Comment").Create(&comment).Error; err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	c.ResSuccess(ctx, nil, "回复成功！")

}
