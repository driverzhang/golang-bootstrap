package page

import (
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/config"
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/orm"
)

var engine *orm.RWEngine

func init() {
	// dsn 根据自己本地或测试服务器数据库地址自行改动，上线时应改为线上服务器数据库地址
	dsn := config.C.Mysql.Bamboo
	var err error
	engine, err = orm.New(dsn, dsn)
	if err != nil {
		panic(err)
	}
}
