package toss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 文件原型是否存在
func IsOriginExist(fileName string) (bool, error) {
	err, path0 := GetOriginPath(fileName)
	if err != nil {
		return false, err
	}
	return GetBucket().IsObjectExist(path0)
}

// 文件是否存在
func IsExist(key string) error {
	isExist, err := cache.Bucket.IsObjectExist(key)
	if !isExist {
		return errors.New("文件不存在:" + key)
	}
	return err
}

// 文件是私有的
func IsPrivate(key string) (bool, error) {
	aclRes, err := cache.Bucket.GetObjectACL(key)
	if err != nil {
		return false, err
	}
	return aclRes.ACL == string(oss.ACLPrivate), err
}

// 文件是开放读
func IsPublicRead(key string) (bool, error) {
	aclRes, err := cache.Bucket.GetObjectACL(key)
	if err != nil {
		return false, err
	}
	return aclRes.ACL == string(oss.ACLPublicRead), err
}

// 文件是开放读写
func IsPublicReadWrite(key string) (bool, error) {
	aclRes, err := cache.Bucket.GetObjectACL(key)
	if err != nil {
		return false, err
	}
	return aclRes.ACL == string(oss.ACLPublicReadWrite), err
}
