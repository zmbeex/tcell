package ttoken

import (
	"encoding/json"
	"errors"
	"github.com/zmbeex/dao/tredis"
	"github.com/zmbeex/gkit"
	"go.uber.org/zap"
	"time"
)

type Token struct {
	Code     string // token 编码
	Id       int64  // 用户ID
	Platform int    // 平台 1androi 2ios 3web
	Time     int64  // 时间戳
	token    string // tokenbase64
}

// @desc 生成token
// @param id tokenID
// @param msg token附加信息
func SetToken(code string, id int64, platform int) (string, error) {
	defer func() {
		r := recover()
		if (r != nil) && (r != "") {
			gkit.Warn("setToken", zap.Any("r", r))
		}
	}()
	m := new(Token)
	m.Id = id
	m.Platform = platform
	m.Code = code
	m.Time = time.Now().Unix()

	if m.Code == "" {
		gkit.Panic("token编码不能为空")
	}
	if m.Platform == 0 {
		gkit.Panic("平台不能为空")
	}
	if m.Id == 0 {
		gkit.Panic("用户ID不能为空")
	}
	// 序列化
	b, err := json.Marshal(m)
	if err != nil {
		return "", errors.New("token序列化失败")
	}

	// 加密
	token, err := gkit.GetSHA(string(b))
	if err != nil {
		return "", errors.New("加密失败")
	}
	// 存储
	err = tredis.SetRedis(token, string(b), time.Duration(Cache.set.ValidTime)*time.Second)
	gkit.CheckPanic(err, "token储存失败")
	gkit.Warn("存储token：" + token)

	return token, nil
}

// @param tokenStr string token密文
func GetToken(tokenStr string) (*Token, error) {
	// 检查token是否重复
	val := tredis.GetRedis("TOKEN#" + tokenStr)
	if len(val) != 0 {
		return nil, errors.New("重复token")
	}
	// 记录token
	err := tredis.SetRedis("TOKEN#"+tokenStr, "token", time.Duration(Cache.set.DiffTime*2)*time.Second)
	if err != nil {
		return nil, err
	}
	token0, err := gkit.GetClientRsa(tokenStr)
	if err != nil {
		return nil, err
	}
	gkit.Warn(token0)
	// 根据token获取用户信息
	tokenInfoStr := tredis.GetRedis(token0)
	token := new(Token)
	err = gkit.GetJson(tokenInfoStr, token)
	if err != nil {
		gkit.Warn("无效token tokenStr：" + tokenStr)
		gkit.Warn("无效token tokenInfoStr：" + tokenInfoStr)
		return nil, errors.New("token已过期")
	}
	return token, nil
}
