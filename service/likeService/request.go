package likeservice

import (
	"fmt"
	"strconv"

	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/gin-gonic/gin"
)

func (c LikeService) ValidataReq(ctx *gin.Context) (model.Like, error) {
	var res model.Like
	owner_id, _ := strconv.Atoi(ctx.Query("ownerid"))
	owner_type := ctx.Query("owner_type")
	selfId := c.GetSelfID(ctx)

	if owner_type != "posts" && owner_type != "comments" {
		return res, fmt.Errorf("不支持此类型的点赞")
	}
	return model.Like{UserID: uint(selfId), OwnerID: uint(owner_id), OwnerType: owner_type}, nil
}
