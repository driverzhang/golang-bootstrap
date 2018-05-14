package main

import (
	_ "git.zhuzi.me/zzjz/zhuzi-bootstrap/router"
	_ "git.zhuzi.me/zzjz/zhuzi-bootstrap/router/page"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/config"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/ginsrv"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/grpcsrv"

	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ginsrv.Listen(config.C.HttpAddr)
		log.Print("http shutdown")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		grpcsrv.Listen(config.C.GrpcAddr)
		log.Print("grpc shutdown")
		wg.Done()
	}()

	wg.Wait()
}
