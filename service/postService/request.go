package postservice

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
	"github.com/PIPIKAI/Ins-gin-vue/server/util/upload"
	"github.com/gin-gonic/gin"
)

type ImgeData struct {
	Uid           int
	Name          string
	Status        string
	Base64Date    string
	Blob          string
	BlobData      interface{}
	Raw           interface{}
	CropedBlobUrl string
	Size          int
	Percentage    int
}

type CreatePostForm struct {
	DataList  []ImgeData `json:"data_list" swaggerignore:"true"`
	Categorys []string   `json:"categorys" example:"example,test"`
	Place     string     `json:"place" example:"school"`
	MyComment string     `json:"mycomment" example:"我的评论"`
	Pictures  []string   `json:"pictures" example:"http://pic.kiass.top/1660055054189wallhaven-rd2jw1_1920x1080.png,http://pic.kiass.top/1660055054194wallhaven-x8eydz.jpg"`
}
type EditPostForm struct {
	Categorys []string `json:"categorys"`
	Place     string   `json:"place"`
	MyComment string   `json:"mycomment"`
}

func (r PostService) UploadImages(dataList []ImgeData) []string {
	res := make([]string, len(dataList))
	for idx, v := range dataList {
		tempf := strings.Split(v.Base64Date, ",")
		file, err := base64.StdEncoding.DecodeString(tempf[len(tempf)-1])
		if err != nil {
			panic("base64.StdEncoding.DecodeString" + err.Error())
		}
		qiniu := upload.GetQiniu()
		log.Println("after GetQiniu")
		url, err := qiniu.Upload(strconv.Itoa(v.Uid)+v.Name, file)
		log.Println("after Upload")

		if err != nil {
			panic("qiniu.Upload" + err.Error())
		}
		res[idx] = url
	}
	return res
}
func (r PostService) CreatePost(ctx *gin.Context) (*CreatePostForm, error) {
	// var v *model.User
	var catch *CreatePostForm
	// 获得绑定的数据
	if err := ctx.ShouldBind(&catch); err != nil {
		log.Println("catch", catch)
		return catch, fmt.Errorf("数据不规范:%v", err.Error())
	}
	//将图片上传 获得url
	if len(catch.Pictures) == 0 {
		log.Println("before UploadImages")
		catch.Pictures = r.UploadImages(catch.DataList)
		log.Println("after UploadImages")

	}

	return catch, nil
}
func (r PostService) EditPost(ctx *gin.Context) (m *EditPostForm, err error) {
	// var v *model.User
	var catch EditPostForm
	if err := ctx.ShouldBind(&catch); err != nil {
		log.Println("catch", catch)
		return nil, fmt.Errorf("数据不规范")
	}
	return &catch, nil
}

func (r PostService) DeletPost(id uint) (m gin.H, err error) {
	v := &model.Post{ID: id}
	if err := common.GetDB().Delete(&v).Error; err != nil {
		return nil, err
	}
	return gin.H{"data": v}, nil
}

func (r PostService) SelectPostById(id uint) (m gin.H, err error) {
	v := &model.Post{ID: id}
	if err := common.GetDB().First(&v).Error; err != nil {
		return nil, err
	}
	return gin.H{"data": v}, nil
}
