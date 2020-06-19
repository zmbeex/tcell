package toss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zmbeex/gkit"
	"path"
)

// 获取软链接指向的原文件
func GetOriginFilePath(symlink string) (error, string) {
	meta, err := GetBucket().GetSymlink(symlink)
	if err != nil {
		return err, ""
	}
	return nil, meta.Get(oss.HTTPHeaderOssSymlinkTarget)
}

// 设置软链接
// symlink 软链接地址
// src 资源地址
func AddSymlink(symlink string, desc string) error {
	err, src := GetOriginPath(symlink)
	if err != nil {
		return err
	}
	// 判断需要备份的文件是否存在
	err = IsExist(src)
	if err != nil {
		return err
	}

	// 创建软链接
	err = GetBucket().PutSymlink(symlink, src)

	// 删除临时文件
	temp := GetTempPath(src)
	err = GetBucket().DeleteObject(temp)
	if err != nil {
		gkit.Warn("删除临时文件失败")
		_ = GetBucket().DeleteObject(symlink)
		return err
	}

	return err
}

// 删除软链接
// path0 软链接地址
func DeleteSymlink(path0 string) error {
	return GetBucket().DeleteObject(path0)
}

// 文件软链接备份
func SymlinkBackUp(path0 string) error {
	// 判断需要备份的文件是否存在
	err := IsExist(path0)
	if err != nil {
		return err
	}
	err = GetBucket().PutSymlink(path.Join(path.Dir(path0), "symlink"+path.Base(path0)), path0)
	return err
}
