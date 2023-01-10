package main

import (
	trippb "coolcar/proto/gen/go"
	"fmt"
	"net"

	"google.golang.org/grpc"
)



func main() {
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	
	trippb.RegisterTripServiceServer(grpcServer,new(trippb.UnimplementedTripServiceServer))
	// 监听端口
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("服务启动失败", err)
		return
	}
	grpcServer.Serve(listen)
	fmt.Println("连接成功")
}