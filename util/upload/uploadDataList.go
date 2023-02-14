package upload

import (
	"encoding/base64"
	"log"
	"strconv"
	"strings"
)

type DataList struct {
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

func UploadImages(dataList []DataList) []string {
	res := make([]string, len(dataList))
	for idx, v := range dataList {
		tempf := strings.Split(v.Base64Date, ",")
		file, err := base64.StdEncoding.DecodeString(tempf[len(tempf)-1])
		if err != nil {
			panic("base64.StdEncoding.DecodeString" + err.Error())
		}
		qiniu := GetQiniu()
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
