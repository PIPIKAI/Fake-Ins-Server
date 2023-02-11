package likeservice

import (
	"fmt"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/gin-gonic/gin"
)

func (c LikeService) IsLiked(req model.Like) bool {
	err := common.GetDB().Model(model.Like{}).First(&req).Error
	return err == nil
}

// @Summary     点赞是否存在
// @Schemes     http https
// @Description LikedOrNot
// @Param       ownerid query string true "ownerid"
// @Param       owner_type query string true "类型"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/likedornot [post]
func (c LikeService) LikedOrNot(ctx *gin.Context) {
	req, err := c.ValidataReq(ctx)
	if err != nil {
		c.ResErr(ctx, "数据验证错误")
		return
	}
	if c.IsLiked(req) {
		c.ResSuccess(ctx, gin.H{"data": true}, "点赞存在")
	} else {
		c.ResSuccess(ctx, gin.H{"data": false}, "点赞不存在")
	}
}

// 点赞一个帖子
// @Summary     点赞一个帖子
// @Schemes     http https
// @Description LikePost
// @Param       ownerid query string true "ownerid"
// @Param       owner_type query string true "类型" example "Post"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/like [post]
func (c LikeService) Like(ctx *gin.Context) {
	req, err := c.ValidataReq(ctx)
	DB := common.GetDB()
	if err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	if req.OwnerType == "posts" {
		if err := DB.First(&model.Post{}, req.OwnerID).Error; err != nil {
			c.ResErr(ctx, "帖子不存在")
			return
		}
		if c.IsLiked(req) {
			c.ResErr(ctx, "已经点赞了")
			return
		}
		if err := DB.Create(&req).Error; err != nil {
			c.ResErr(ctx, "创建点赞失败")
			return
		}

		if err := DB.Model(&model.Post{ID: req.OwnerID}).Association("Likes").Append(&req); err != nil {
			c.ResErr(ctx, err.Error())
			return
		}
	} else if req.OwnerType == "comments" {
		if err := DB.First(&model.Comment{}, req.OwnerID).Error; err != nil {
			c.ResErr(ctx, "评论不存在")
			return
		}
		if err := DB.Create(&req).Error; err != nil {
			c.ResErr(ctx, "创建点赞失败")
			return
		}
		if err := DB.Model(&model.Comment{ID: req.OwnerID}).Association("Likes").Append(&req); err != nil {
			c.ResErr(ctx, err.Error())
			return
		}
	} else {
		c.ResErr(ctx, "不存在该类型")
		return
	}
	c.ResSuccess(ctx, nil, "点赞成功")
}

// 取消点赞一个帖子
// @Summary     取消点赞一个帖子
// @Schemes     http https
// @Description UndoLike
// @Param       ownerid query int true "ownerid"
// @Param       owner_type query string true "owner_type"
// @Tags        Post
// @Accept      json
// @Produce     json
// @success     200
// @Router      /post/undolike [delete]
func (c LikeService) UndoLike(ctx *gin.Context) {
	req, err := c.ValidataReq(ctx)
	DB := common.GetDB()
	if err != nil {
		c.ResErr(ctx, err.Error())
		return
	}
	if req.OwnerType == "posts" {
		if err := DB.First(&req).Error; err != nil {
			c.ResErr(ctx, "没有点赞")
			return
		}
		if err := DB.Delete(&req).Error; err != nil {
			c.ResErr(ctx, "删除点赞失败")
			return
		}
		if err := DB.Model(&model.Post{ID: req.OwnerID}).Association("Likes").Delete(&req); err != nil {
			c.ResErr(ctx, err.Error())
			return
		}
	} else if req.OwnerType == "comments" {
		// if err := DB.First(&model.Comment{}, req.OwnerID).Error; err != nil {
		// 	c.ResErr(ctx, "评论不存在")
		// 	return
		// }
		// if err := DB.Model(&model.Comment{ID: req.OwnerID}).Association("Likes").Delete(&req); err != nil {
		// 	c.ResErr(ctx, err.Error())
		// 	return
		// }
		like := &model.Like{
			UserID:    req.UserID,
			OwnerID:   req.OwnerID,
			OwnerType: req.OwnerType,
		}
		if err := DB.First(&like).Error; err != nil {
			c.ResErr(ctx, "点赞不存在")
			return
		}
		fmt.Println("like:", like)

		if err := DB.Preload("User").Preload("Post").Preload("Comment").Preload("Like").Delete(&like).Error; err != nil {
			c.ResErr(ctx, err.Error())
			return
		}
	} else {
		c.ResErr(ctx, "不存在该类型")
		return
	}
	c.ResSuccess(ctx, nil, "取消点赞成功")
}

func (c LikeService) DisLike(ctx *gin.Context) {

}
func (c LikeService) UndoDisLike(ctx *gin.Context) {

}
