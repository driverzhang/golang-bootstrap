package model

import (
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/grpcclt"
)

// 实际在编写model层的时候，应该在model包里init这个client
var client grpcclt.Client

func init() {
	dsn := ":8081"
	var err error
	client, err = grpcclt.New(dsn)
	if err != nil {
		panic(err)
	}
}
