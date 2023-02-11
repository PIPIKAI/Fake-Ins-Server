package upload

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/PIPIKAI/Ins-gin-vue/server/util"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
)

type Qiniu struct {
	domain       string
	accessKey    string
	secreKey     string
	bucket       string
	upToken      string
	formUploader *storage.FormUploader
}

var qiniu *Qiniu

func NewQiniuUper() {
	qiniu = &Qiniu{
		domain:    viper.GetString("qiniu.domain"),
		accessKey: viper.GetString("qiniu.accessKey"),
		secreKey:  viper.GetString("qiniu.secreKey"),
		bucket:    viper.GetString("qiniu.bucket"),
	}
	qiniu.initialize()
}
func GetQiniu() *Qiniu {
	qiniu.initialize()
	return qiniu
}
func (q *Qiniu) initialize() {
	accessKey := q.accessKey
	secretKey := q.secreKey
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := q.bucket

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}

	q.upToken = putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	q.formUploader = storage.NewFormUploader(&cfg)
}
func (q *Qiniu) Upload(filename string, file []byte) (string, error) {
	ret := storage.PutRet{}
	err := q.formUploader.Put(context.Background(), &ret, q.upToken, filename, bytes.NewReader(file), int64(len(file)), nil)
	if err != nil {
		return "", err
	}
	return q.publicAccessURL(filename), nil
}
func (q *Qiniu) publicAccessURL(key string) string {
	return storage.MakePublicURL(q.domain, key)
}

type FileForm struct {
	Name string `json:"name" binding:"required"`
	File string `json:"file" binding:"required"`
}

type UploadController struct {
}

func NewUploadController() UploadController {
	return UploadController{}
}

func (c UploadController) UploadImages(ctx *gin.Context) {
	var catch *FileForm
	if err := ctx.ShouldBind(&catch); err != nil {
		util.Response.Error(ctx, nil, "数据不规范")
		return
	}
	// tempf := strings.Split(catch.File, ",")
	// catch.File = tempf[len(tempf)-1]
	file, err := base64.StdEncoding.DecodeString(catch.File)
	if err != nil {
		util.Response.Error(ctx, nil, err.Error())
		return
	}
	qiniu := GetQiniu()
	url, err := qiniu.Upload(catch.Name, file)
	if err != nil {
		util.Response.Error(ctx, nil, err.Error())
		return
	}
	util.Response.Success(ctx, gin.H{"url": url}, "200")
}
