package model

import (
	"context"

	"git.zhuzi.me/zzjz/zhuzi-payment/pb"
)

func GetUserName() (string, error) {
	c := pb.NewDataClient(client)
	u, err := c.GetUser(context.Background(), &pb.UserRq{Id: 1})
	return u.Name, err
}
