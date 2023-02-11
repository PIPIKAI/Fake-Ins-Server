package util

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Unix2String(now int64) string {
	return strconv.Itoa(int(now))
}
func NowUnix() string {
	return Unix2String(time.Now().Unix())
}
func RandomStr(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm123456790QWERTYUIOPASDFGHJKLZXCVBNM")
	var res string
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		res += string(letters[rand.Int()%len(letters)])
	}
	return res
}

func PagePreSolve(ctx *gin.Context) (int, int) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 15
	case pageSize <= 0:
		pageSize = 6
	}
	return page, pageSize
}
