package toss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zmbeex/gkit"
	"path"
	"strings"
)

type tossSetting struct {
	Endpoint        string `title:"资源域名"`
	AccessKeyId     string `title:"验证键" defaultValue:"AccessKeyId"`
	AccessKeySecret string `title:"验证值" defaultValue:"AccessKeySecret"`
	BucketName      string `title:"桶" defaultValue:"BucketName"`
	Referer         string `title:"防盗链xxx,xxx,xxx"`

	Origin string `title:"md5目录，所有文件都存储在这个文件夹" defaultValue:"z/md5"`
	User   string `title:"用户资源目录" defaultValue:"z/u/"`
	Temp   string `title:"临时资源目录" defaultValue:"z/t/"`
	Group  string `title:"团队资源目录" defaultValue:"z/g/"`
}

var cache struct {
	Set    *tossSetting
	Bucket *oss.Bucket
	Client *oss.Client
}

func init() {
	cache.Set = new(tossSetting)
	gkit.InitSetting("toss", cache.Set, "oss配置", func() {
		// 创建OSSClient实例。
		client, err := oss.New(cache.Set.Endpoint, cache.Set.AccessKeyId, cache.Set.AccessKeySecret)
		gkit.CheckPanic(err, "创建连接异常")
		cache.Client = client

		// oss 文件桶
		bucket, err := client.Bucket(cache.Set.BucketName)
		gkit.CheckPanic(err, "获取存储空间异常")
		cache.Bucket = bucket
	})
}

// 获取文件存储路径
// fileName 文件名 md5+后缀
func GetOriginPath(fileName string) (error, string) {
	if len(path.Base(fileName)) < 32 {
		return errors.New("无效文件名: " + fileName), ""
	}
	s := cache.Set.Origin + "/" + path.Base(fileName)
	s = strings.ReplaceAll(s, "//", "/")
	return nil, s
}

// 用户文件目录
func GetUserPath(userId int64, keys ...string) string {
	s := cache.Set.User + "/" + gkit.ToString(userId) + "/" + strings.Join(keys, "/")
	s = strings.ReplaceAll(s, "//", "/")
	s = strings.Trim(s, "/")
	return s
}

// 临时文件目录
// srcPath 资源临时文件目录
func GetTempPath(srcPath string) string {
	s := cache.Set.Temp + "/" + path.Base(srcPath)
	s = strings.ReplaceAll(s, "//", "/")
	s = strings.Trim(s, "/")
	return s
}

// 团队目录
// 分组 group
// 团队id groupId
func GetGroupPath(group string, groupId interface{}, keys ...string) string {
	s := cache.Set.Group + "/" + group + "/" + gkit.ToString(groupId) + "/" + strings.Join(keys, "/")
	s = strings.ReplaceAll(s, "//", "/")
	s = strings.Trim(s, "/")
	return s
}

func GetBucket() *oss.Bucket {
	return cache.Bucket
}

//type Oss struct {
//	option []oss.Option
//	bucket *oss.Bucket
//}
//
//
//func (o *Oss) SetBucketName(name string, referees string) (*oss.Bucket, error) {
//	// 判断存储空间是否存在。
//	isExist, err := cache.Client.IsBucketExist(name)
//	gkit.CheckPanic(err, "连接存储空间异常")
//	if !isExist {
//		// 创建存储空间，并设置数据容灾类型为同城区域冗余存储。
//		err = cache.Client.CreateBucket(name, oss.RedundancyType(oss.RedundancyZRS))
//		if err != nil {
//			gkit.Warn("创建存储空间异常", zap.String("bucketName", name), zap.String("err", err.Error()))
//			return nil, err
//		}
//	}
//	if cache.Set.Referer != "" {
//		str := strings.ReplaceAll(referees, " ", "")
//		if str == "" {
//			str = cache.Set.Referer
//		}
//		referees := strings.Split(str, ",")
//		err = cache.Client.SetBucketReferer(name, referees, false)
//		if err != nil {
//			gkit.Warn("设置盗链", zap.String("bucketName", name), zap.String("err", err.Error()))
//			return nil, err
//		}
//	}
//
//	// 获取存储空间。
//	bucket, err := cache.Client.Bucket(name)
//	if err != nil {
//		gkit.Warn("获取存储空间异常", zap.String("bucketName", name), zap.String("err", err.Error()))
//		return nil, err
//	}
//	return bucket, err
//}
//
//// 设置一个操作
//func (o *Oss) SetOption(option oss.Option) {
//	o.option = append(o.option, option)
//}
//
//// 定义进度条监听器。
//type ossProgressListener struct {
//	handle func(event *oss.ProgressEvent)
//}
//
//// 定义进度变更事件处理函数。
//func (listener *ossProgressListener) ProgressChanged(event *oss.ProgressEvent) {
//	listener.handle(event)
//}
//
//// 设置进度条方法
//func (o *Oss) CreateProgress(f func(event *oss.ProgressEvent)) {
//	if f != nil {
//		listener := &ossProgressListener{}
//		listener.handle = func(event *oss.ProgressEvent) {
//			f(event)
//		}
//		o.SetOption(oss.Progress(listener))
//	}
//}
//
//// 设置压缩
//func (o *Oss) SetGzip() {
//	o.SetOption(oss.AcceptEncoding("gzip"))
//}
//
//// 设置文件过期时间
//// 文件自动设置为公共只读文件
//func (o *Oss) SetExpires(t time.Duration) {
//	d := time.Date(2020, time.April, 10, 23, 0, 0, 0, time.UTC)
//	d.Add(t)
//	o.option = append(o.option, oss.Expires(d))
//	o.option = append(o.option, oss.ObjectACL(oss.ACLPublicRead))
//}
//
//// 设置meta X-Oss-Meta-xxx
//func (o *Oss) SetMete(key string, value string) {
//	o.option = append(o.option, oss.Meta(key, value))
//}
//
//// 禁止覆盖
//func (o *Oss) ForbidOverWrite() {
//	// 禁止覆盖同名文件
//	o.option = append(o.option, oss.ForbidOverWrite(true))
//}
