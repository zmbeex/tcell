package tossDogApi

import (
	"github.com/zmbeex/dog"
	"github.com/zmbeex/gkit"
	"github.com/zmbeex/tcell/toss"
)

type ossUploadBase64Params struct {
	Base64  string `title:"base64文件" check:"notNull"`
	Md5Name string `title:"文件内容的md5值" check:"notNull"`
}

func ossUploadBase64(d *dog.Dog) {
	params := new(ossUploadBase64Params)
	d.GetParams(params)

	data := gkit.Base64ToString(params.Base64)
	if len(data) == 0 {
		gkit.Panic("base64数据异常")
	}
	err := toss.UploadFile(params.Md5Name, []byte(data))
	gkit.CheckPanic(err, "上传文件失败")
}
