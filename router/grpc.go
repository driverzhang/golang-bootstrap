package route

import (
	"git.zhuzi.me/zzjz/zhuzi-payment/api/agrpc"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/grpcsrv"
	"git.zhuzi.me/zzjz/zhuzi-payment/pb"
)

// 注册grpc
func init() {
	pb.RegisterDataServer(grpcsrv.Server, &agrpc.UserService{})
}
