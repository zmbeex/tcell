package tfile

import (
	"encoding/json"
	"errors"
	"github.com/zmbeex/gkit"
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// 当前命令备份历史
var backupHistoryList = make(map[string]bool)

type backupInfoModal struct {
	Sha1 string // 备份文件sha1
	Path string // 备份文件路径
}

// 整体备份信息
var backupInfo = make(map[string]map[string]string)

var backupWriteList = []string{
	string(os.PathSeparator) + "tmp",
	string(os.PathSeparator) + "backup",
}

// @desc 设置备份信息
func BackUpSafe(path0 string) {
	// 路由矫正
	path1 := gkit.FixPathToAbs(path0)
	for _, writeS := range backupWriteList {
		if gkit.IsStringIncludeString(path1, writeS) {
			gkit.Debug("白名单路径，不备份", zap.String("path", path1))
			return
		}
	}

	gkit.Debug("备份目录", zap.String("path", path1))

	if gkit.IsExisted(path1).IsFile() {
		backupFile(path1)
	} else if gkit.IsExisted(path1).IsDir() {
		//pathList, err := gkit.WalkDir(path1, "")
		//		//gkit.CheckPanic(err, "备份文件夹异常")
		//		//for _, path3 := range pathList {
		//		//	backupFile(path3)
		//		//}
	}
}

// 备份文件
func backupFile(path0 string) {
	// 读取文件
	fileS := gkit.ReadFile(path0)
	// 检查文件当前命令是否已备份过
	if backupHistoryList[path0] {
		gkit.Debug("非初始文件，无需备份")
		return
	}
	gkit.Info("备份文件", zap.String("path", path0))
	// 获取文件sha1
	sha1, err := gkit.GetMd5(string(fileS))
	if err != nil {
		gkit.Warn(err.Error())
	}
	// 备份文件目录
	bcpPath := gkit.GetWorkspace("backup/file/" + sha1 + ".bcp")
	// 检查是否已存在
	if gkit.IsExisted(bcpPath).IsFile() {
		gkit.Debug("文件已存在，无需备份")
		return
	}
	// 写入文件到备份文件夹
	err = WriteFileSafe(gkit.FixPath(bcpPath), fileS)
	gkit.CheckPanic(err, "文件备份写入失败")

	backupHistoryList[path0] = true

	// 备份信息文件路径
	bkInfoPath := gkit.GetWorkspace("backup/backupInfo.json")
	_ = gkit.MakeDir(path.Dir(bkInfoPath))
	// 读取备份或创建信息文件
	if len(backupInfo) == 0 {
		bkInfoByte := gkit.ReadFile(bkInfoPath)
		if len(bkInfoByte) == 0 {
			backupInfo = map[string]map[string]string{}
		} else {
			err := gkit.GetJson(string(bkInfoByte), &backupInfo)
			gkit.CheckPanic(err, "读取备份配置信息失败")
		}
	}
	// 写入信息文件
	if len(backupInfo[gkit.ToString(gkit.SysVal.AppStartTime)]) == 0 {
		backupInfo[gkit.ToString(gkit.SysVal.AppStartTime)] = make(map[string]string)
	}
	// 添加信息
	backupInfo[gkit.ToString(gkit.SysVal.AppStartTime)][sha1] = path0
	// 持久化
	err = gkit.WriteFile(gkit.FixPath(bkInfoPath), []byte(gkit.SetJson(backupInfo)))
	gkit.CheckPanic(err, "持久化备份信息失败")
}

// @desc 获取备份信息
func GetBackupInfo(time string) (map[string]string, error) {
	// 读取备份目录所有文件夹
	list, err := gkit.ListDirDir(gkit.FixPathToAbs(Cache.Set.BackupPath), "")
	if err != nil {
		return nil, err
	}
	vaildName := ""
	for _, val := range list {
		name := filepath.Base(val)
		reg, err := regexp.Compile(`[^0-9]`)
		if err != nil {
			return nil, err
		}
		name = reg.ReplaceAllString(name, "")

		newName := strings.Replace(name, time, "", -1)
		if vaildName == newName {
			return nil, errors.New("备份又多条符合要求")
		}
		if newName != name {
			vaildName = newName
		}
	}

	// 读取配置文件
	jsonB := gkit.ReadFile(gkit.FixPath(Cache.Set.BackupPath + "/" + vaildName + "/backupInfo.json"))
	jsonMap := make(map[string]string)
	err = json.Unmarshal(jsonB, jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}
