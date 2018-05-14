package grpcsrv

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

var Server *grpc.Server

func Listen(addr string) {
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		GracefulStop()
	}()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = Server.Serve(lis)
	if err != nil {
		log.Printf("grpc Serve err: %v", err)
	}
}

func GracefulStop() {
	Server.GracefulStop()
}

func init() {
	Server = grpc.NewServer()
}
