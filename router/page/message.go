package page

import (
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/api/page/message"
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/ginsrv"
)

func init() {
	apiV1 := ginsrv.Engine.Group(`/api/v2`)

	apiV1.DELETE(`/sites/:siteid/page/:pageid`, message.Delete) //删除页面

}
