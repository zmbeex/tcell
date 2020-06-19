package tossDogApi

import (
	"github.com/zmbeex/dog"
	"github.com/zmbeex/tcell/toss"
	"net/http"
	"time"
)

type ossDownloadParams struct {
	FileName string `title:"文件名称" check:"notNull"`
}

func ossDownload(d *dog.Dog) {
	params := new(ossDownloadParams)
	d.GetParams(params)
	err, data := toss.DownloadFileToBytes(params.FileName)
	_, err = d.W.Write(data)
	if err != nil {
		panic(404)
	}
	// 文件缓存30天
	d.W.Header().Set("Expires", time.Now().Add(time.Hour*24*30).UTC().Format(http.TimeFormat))
	panic(200)
}
