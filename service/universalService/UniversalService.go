package universalservice

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-gonic/gin"
)

func (c UniversalService) GetSelf(ctx *gin.Context) model.User {
	tp, _ := ctx.Get("user")
	return tp.(model.User)
}
func (c UniversalService) GetSelfID(ctx *gin.Context) uint {
	return c.GetSelf(ctx).ID
}
func (c UniversalService) GetSelfUserName(ctx *gin.Context) string {
	return c.GetSelf(ctx).UserName
}

func (c UniversalService) ResErr(ctx *gin.Context, msg string) {
	util.Response.Error(ctx, nil, msg)
}
func (c UniversalService) ResSuccess(ctx *gin.Context, h gin.H, msg string) {
	util.Response.Success(ctx, h, msg)
}
