package main

import (
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripService"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)



func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":8081")
	if(err != nil) {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	
	trippb.RegisterTripServiceServer(grpcServer, &trip.Service{})
	
	// if err != nil {
	// 	fmt.Println("服务启动失败", err)
	// 	return
	// }
	log.Fatal(grpcServer.Serve(lis))
	fmt.Println("连接成功")
}