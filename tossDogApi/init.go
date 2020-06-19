package tossDogApi

import (
	"github.com/zmbeex/dog"
)

func Run() {
	dog.POST("", "/oss/ossUpload", "上传文件oss", nil, ossUpload)
	dog.POST("", "/oss/ossUploadBase64", "base64格式上传文件", ossUploadBase64Params{}, ossUploadBase64)
	dog.POST("", "/oss/ossDownload", "下载文件oss", ossDownloadParams{}, ossDownload)
	dog.POST("", "/oss/ossIsExist", "检查文件是否存在oss", ossIsExistParams{}, ossIsExist)
}
