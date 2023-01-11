package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripService"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)



func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
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
func startGRPCGateway() {
	// 建立空的上下文
	c := context.Background();
	c, cancel := context.WithCancel(c);
	defer cancel()
	// 分发器，当有服务的时候，分发到哪个服务器上面
	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
	if(err != nil) {
		log.Fatalf("cannot start grpc gateway %v", err)
	}
	// 8080的请求分发到8081的服务器上面
	err = http.ListenAndServe(":8080", mux)
	if( err != nil) {
		log.Fatalf("cannot listen to 8080 %v", err)
	}
}

