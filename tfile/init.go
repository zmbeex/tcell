package tfile

type Set struct {
	BackupPath string `title:"操作备份路径, 配置有效路径后开启备份" defaultValue:""`
}

var Cache struct {
	Set *Set
}
