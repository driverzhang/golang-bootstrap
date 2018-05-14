package agrpc

import (
	"git.zhuzi.me/zzjz/zhuzi-payment/pb"
	"golang.org/x/net/context"
)

type UserService struct {
}

func (*UserService) GetUser(context.Context, *pb.UserRq) (*pb.UserRp, error) {
	return &pb.UserRp{Name: "bysir"}, nil
}
