package ttoken

import (
	"github.com/zmbeex/gkit"
)

type ttokenSetting struct {
	ValidTime int   `title:"token有效期(默认一天)" defaultValue:"86400"`
	DiffTime  int64 `title:"时间差,秒" check:"notNull" defaultValue:"300"`
}

var Cache struct {
	set *ttokenSetting
}

func init() {
	Cache.set = new(ttokenSetting)
	gkit.InitSetting("ttoken", Cache.set, "token", func() {

	})
}
