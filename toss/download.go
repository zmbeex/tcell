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

//import (
//	"github.com/aliyun/aliyun-oss-go-sdk/oss"
//	"io/ioutil"
//)
//
//// 下载文件的字节流
//func (o *Oss) DownloadBytes(key string) []byte {
//	body, err := cache.Bucket.GetObject(key, o.option...)
//	if err != nil {
//		return []byte{}
//	}
//	defer func() {
//		_ = body.Close()
//	}()
//	data, err := ioutil.ReadAll(body)
//	if err != nil {
//		return []byte{}
//	}
//	return data
//}
//
//// 下载文件到本地
//// 键
//// 下载到本地的路径
//func (o *Oss) DownloadLocal(key string, localPath string) error {
//	err := cache.Bucket.GetObjectToFile(key,
//		localPath,
//		o.option...)
//	return err
//}
//
//// 断点续下载
//func (o *Oss) DownloadBreakpointRenewal(key string, localPath string) error {
//	var ops = o.option
//	// 使用协程 3个
//	ops = append(ops, oss.Routines(3))
//	// 开启断点续传
//	ops = append(ops, oss.Checkpoint(true, ""))
//	err := cache.Bucket.DownloadFile(key,
//		localPath,
//		100*1024,
//		ops...)
//	return err
//}
