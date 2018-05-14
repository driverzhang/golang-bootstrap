package grpcclt

import (
	"google.golang.org/grpc"
)

type Client *grpc.ClientConn

func New(addr string) (Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
