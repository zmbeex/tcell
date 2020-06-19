package tossDogApi

import (
	"github.com/zmbeex/dog"
	"github.com/zmbeex/gkit"
	"github.com/zmbeex/tcell/toss"
	"strings"
)

type ossIsExistParams struct {
	FileList string `title:"文件列表 xxx,yyy,zzz"`
}

func ossIsExist(d *dog.Dog) {
	params := new(ossIsExistParams)
	d.GetParams(params)
	params.FileList = strings.ReplaceAll(params.FileList, " ", "")
	files := strings.Split(params.FileList, ",")
	var list []string
	for _, filePath := range files {
		err, p := toss.GetOriginPath(filePath)
		gkit.CheckPanic(err, "文件名存在错误1")
		ok, err := toss.IsOriginExist(p)
		gkit.CheckPanic(err, "文件名存在错误2")
		if !ok {
			list = append(list, filePath)
		}

	}
	d.SetResult(&list)
}
