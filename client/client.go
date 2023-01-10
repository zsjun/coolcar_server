package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	// 建立链接
	// grpc.WithInsecure()
	conn, err := grpc.Dial("127.0.0.1:8081",grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial Error ", err)
		return
	}
	// 延迟关闭链接
	defer conn.Close()
	// 实例化客户端
	client := trippb.NewTripServiceClient(conn)
	// 发起请求
	reply, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip456"})
	if err != nil {
		fmt.Println("启动失败", err)
		return
	}
	fmt.Println("返回:", reply)
}