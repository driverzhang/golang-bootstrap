package grpcclt

import (
	"context"
	"log"
	"testing"
	"time"

	"git.zhuzi.me/zzjz/zhuzi-payment/pb"
)

// 请先启动GrpcService后测试
func TestGrpcUser(t *testing.T) {
	for {
		c := pb.NewDataClient(client)
		u, err := c.GetUser(context.Background(), &pb.UserRq{Id: 1})
		if err != nil {
			log.Printf("%+v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Printf("%+v", u)

		time.Sleep(1 * time.Second)
	}
}

// 实际在编写model层的时候，应该在model包里init这个client
// 和mysql一样, 属于数据层, 不应该在除了model里的其他地方调用
var client Client

func init() {
	dsn := ":8081"
	var err error
	client, err = New(dsn)
	if err != nil {
		panic(err)
	}
}
