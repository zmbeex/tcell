package tossDogApi

import (
	"github.com/zmbeex/dog"
	"github.com/zmbeex/gkit"
	"github.com/zmbeex/tcell/toss"
	"io"
	"os"
	"path"
)

type fileInfo struct {
	FileName string
	FileType string
	Md5      string
}

func ossUpload(d *dog.Dog) {
	// 无论用的什么路由，原理是要从request获取数据
	r := d.R
	// request 获得文件 reader
	reader, err := r.MultipartReader()
	gkit.CheckPanic(err, "1")
	if reader == nil {
		gkit.CheckPanic(err, "2")
	}
	var list []*fileInfo

	//遍历操作 获得的
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if part == nil {
			continue
		}
		if part.FileName() == "" {
			gkit.Warn("存在无效文件")
			continue
		} else {
			func() {
				//创建临时文件
				tempFile := gkit.GetWorkspace("temp/" + part.FileName())
				_ = gkit.MakeDir(path.Dir(tempFile))

				dst, err := os.Create(tempFile)
				gkit.CheckPanic(err, "3")
				defer func() {
					dst.Close()
					err := os.Remove(tempFile)
					if err != nil {
						gkit.Warn("删除临时文件：" + tempFile + "失败")
					}
				}()

				//将获取到的文件复制 给 创建的文件
				_, err = io.Copy(dst, part)
				if err != nil {
					gkit.CheckPanic(err, "4")
				}

				// 获取md5值
				fileBytes := gkit.ReadFile(tempFile)
				md5Str, err := gkit.GetMd5(string(fileBytes))
				gkit.CheckPanic(err, "md5")

				// 返回结果
				info := new(fileInfo)
				err, fileName := toss.GetOriginPath(md5Str + path.Ext(part.FileName()))
				gkit.CheckPanic(err, "获取路径失败")
				info.FileName = fileName
				info.FileType = part.Header.Get("type")
				info.Md5 = md5Str
				list = append(list, info)

				// 上传到oss
				err = toss.UploadFile(info.FileName, fileBytes)
				gkit.CheckPanic(err, "文件上传失败2")
			}()
		}
	}

	d.SetResult(&list)
}
