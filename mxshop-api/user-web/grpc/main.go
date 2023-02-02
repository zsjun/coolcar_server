package main

import (
	"flag"
	"fmt"
	handler "mxshop-api/user-web/grpc/handler"
	"mxshop-api/user-web/grpc/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server,&handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)

	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}