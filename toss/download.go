package toss

import "io/ioutil"

// 下载文件
func DownloadFileToBytes(symlink string) (error, []byte) {
	// 根据软链接查询原文件
	err, path0 := GetOriginFilePath(symlink)
	if err != nil {
		return err, []byte{}
	}
	// 读取文件
	body, err := GetBucket().GetObject(path0)
	if err != nil {
		return err, []byte{}
	}
	defer func() {
		_ = body.Close()
	}()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err, []byte{}
	}
	return nil, data
}
