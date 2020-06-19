package toss

import (
	"bytes"
	"errors"
)

// 文件上传
func UploadFile(path0 string, bs []byte) error {
	err, filePath := GetOriginPath(path0)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return errors.New("不支持保存空文件")
	}

	// 创建临时软链接
	err = GetBucket().PutSymlink(GetTempPath(filePath), filePath)
	if err != nil {
		return err
	}

	// 上传文件
	err = GetBucket().PutObject(filePath, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	return nil
}

//// 设置为可用文件
//func SetToUserFile(srcPath string, userId int64, keyword ...string) (error, string) {
//	// user
//	src := "u/" + gkit.ToString(userId) + "/"
//	if len(keyword) > 0 {
//		src += strings.Join(keyword, "/") + "/"
//	}
//	src += path.Base(srcPath)
//
//	_, err := GetBucket().CopyObject(srcPath, src)
//	return err, src
//}
//
//// 设置为可用文件
//func SetToFile(srcPath string, keyword ...string) (error, string) {
//	// common
//	src := "c/"
//	if len(keyword) > 0 {
//		src += strings.Join(keyword, "/") + "/"
//	}
//	src += strings.TrimPrefix(srcPath, "md5")
//	_, err := GetBucket().CopyObject(srcPath, src)
//	return err, src
//}
//// 上传文件，通过字节
//func (o *Oss) UploadBytes(key string, fileBytes []byte) error {
//	return cache.Bucket.PutObject(key, bytes.NewReader(fileBytes), o.option...)
//}
//
//// 上次文件，通过文件文件地址
//func (o *Oss) UploadLocal(key string, filePath string) error {
//	err := cache.Bucket.PutObjectFromFile(key, filePath, o.option...)
//	return err
//}
//
//// 断点续传
//func (o *Oss) UploadBreakpointRenewal(key string, filePath string) error {
//	var ops = o.option
//	// 使用协程 3个
//	ops = append(ops, oss.Routines(3))
//	// 开启断点续传
//	ops = append(ops, oss.Checkpoint(true, ""))
//
//	// key 存储键
//	err := cache.Bucket.UploadFile(key,
//		// filePath 文件路径
//		filePath,
//		// 分片大小 100k
//		100*1024,
//		o.option...)
//	return err
//}
