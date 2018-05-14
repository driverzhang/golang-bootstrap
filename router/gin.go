package route

import (
	"git.zhuzi.me/zzjz/zhuzi-payment/api/agin"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/ginsrv"
)

// 注册gin
func init() {
	v1 := ginsrv.Engine.Group("/api/v1")

	v1.GET("/user", agin.GetUser)
}
