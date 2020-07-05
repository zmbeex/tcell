package tfile

import (
	"github.com/zmbeex/gkit"
	"go.uber.org/zap"
	"path"
)

// @desc 创建文件, 安全模式
// @desc 有则创建，没有则覆盖
// @param targetPath 文件路径
// @param content 文件内容
func WriteFileSafe(targetPath string, content []byte) error {
	// 路由修正
	path0 := gkit.FixPathToAbs(targetPath)
	gkit.Debug("WriteFileSafe", zap.String("path", path0))

	s := gkit.ReadFile(path0)
	if gkit.FilterBlankString(string(content)) == gkit.FilterBlankString(string(s)) {
		return nil
	}

	// 文件夹补全
	_ = gkit.MakeDir(path.Dir(path0))

	// 备份
	//BackUpSafe(path0)
	// 写入
	return gkit.WriteFile(path0, content)
}

// @desc 读取文件
// @path 文件路径
// @return 返回文件字符串
func ReadFileSafe(path0 string) []byte {
	gkit.Debug("ReadFileSafe", zap.String("path", path0))
	// 路由修正
	path := gkit.FixPathToAbs(path0)

	// 读取
	return gkit.ReadFile(path)
}

// @desc 逐行读取文件
// @param 读取文件路径
// @operate 操作
func ReadFileLineSafe(path0 string, operate func(lineStr string)) error {
	gkit.Debug("ReadFileLineSafe", zap.String("path", path0))
	// 路径修正
	path := gkit.FixPathToAbs(path0)

	// 读取
	return gkit.ReadFileLine(path, operate)
}

// @desc 复制文件
// @param src资源目录
// @param des目标目录
func CopyFileSafe(src string, des string) error {
	// 路径修正
	srcPath := gkit.FixPathToAbs(src)
	desPath := gkit.FixPathToAbs(des)
	gkit.Debug("CopyFileSafe", zap.String("src", srcPath), zap.String("des", desPath))

	s := gkit.ReadFile(srcPath)
	d := gkit.ReadFile(desPath)
	if gkit.FilterBlankString(string(d)) == gkit.FilterBlankString(string(s)) {
		return nil
	}

	// 文件夹补全
	err := gkit.MakeDir(path.Dir(desPath))
	if err != nil {
		return err
	}

	// 备份
	//BackUpSafe(desPath)

	// 复制
	return gkit.CopyFile(srcPath, desPath)
}

// @desc 复制文件夹
// @param 复制资源目录
// @param 目录目录
func CopyDirSafe(srcPath string, desPath string) error {
	gkit.Debug("CopyDirSafe", zap.String("src", srcPath), zap.String("des", desPath))
	// 路径修正
	src := gkit.FixPathToAbs(srcPath)
	des := gkit.FixPathToAbs(desPath)

	// 文件夹补全
	err := gkit.MakeDir(des)
	if err != nil {
		return err
	}

	// 备份
	//BackUpSafe(des)

	// 复制
	return gkit.CopyDir(src, des)
}

// @desc 删除文件
// @param 删除文件目录
func RemoveFileSafe(path0 string) error {
	// 路径修正
	path1 := gkit.FixPathToAbs(path0)

	gkit.Error("RemoveFileSafe: " + path1)
	// 备份
	//BackUpSafe(path.Dir(path1))

	// 删除
	return gkit.RemoveFile(path1)
}

// @desc 删除文件夹
func RemoveDirSafe(path0 string) error {
	gkit.Debug("RemoveDirSafe", zap.String("path", path0))
	if !gkit.IsExisted(path0).IsDir() {
		return nil
	}
	// 路径修正
	path1 := gkit.FixPathToAbs(path0)

	// 备份
	//BackUpSafe(path1)

	// 删除
	return gkit.RemoveDir(path1)
}

// @desc 删除文件
// @param 删除文件目录
func RenameFileSafe(path0 string, target0 string) error {
	gkit.Debug("RenameFileSafe", zap.String("path", path0), zap.String("to", target0))
	// 路径修正
	path1 := gkit.FixPathToAbs(path0)
	target := gkit.FixPathToAbs(target0)

	// 文件夹补全
	err := gkit.MakeDir(path.Dir(target))
	if err != nil {
		return err
	}

	// 备份
	//BackUpSafe(path.Dir(path1))
	//BackUpSafe(path.Dir(target))

	// 重命名
	return gkit.RenameFile(path1, target)
}
