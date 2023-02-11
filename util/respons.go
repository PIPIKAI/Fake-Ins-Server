package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseStruct struct{}

var Response ResponseStruct

func InitResponseStruct() {
	Response = ResponseStruct{}
}

func (ResponseStruct) ResponsFmt(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}
func (r ResponseStruct) Success(ctx *gin.Context, data gin.H, msg string) {
	r.ResponsFmt(ctx, http.StatusOK, 200, data, msg)
}
func (r ResponseStruct) Error(ctx *gin.Context, data gin.H, msg string) {
	r.ResponsFmt(ctx, http.StatusBadRequest, 400, data, msg)
}
