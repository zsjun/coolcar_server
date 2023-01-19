package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {


	interceptor := func(ctx context.Context,method string, req, reply interface{}, cc *grpc.ClientConn,invoke grpc.UnaryInvoker, opts ...grpc.CallOption)(error) {
		start := time.Now();
		err := invoke(ctx, method,req,reply,cc,opts...);
		fmt.Println(time.Since(start));
		return err;
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	// grpc.u
	// opt := grpc.UnaryInterceptor(interceptor)

	conn, err := grpc.Dial("localhost:50052",grpc.WithInsecure(),opt);

	if err != nil {
		panic(err);
	}
	defer conn.Close();

	// c := streampb.NewGreeterClient(conn);

	// // md := metadata.Pairs("timestamp", time.Now().Format("10000"));
	// md := metadata.New(
	// 	map[string]string{
	// 		"name":"bobby",
	// 		"password": "sdd",
	// 	},
	// )
	// ctx := metadata.NewOutgoingContext(context.Background(),md)
	// res, _ := c.GetStream(context.Background(),&streampb.StreamReq{Data:"慕课网"});

	// for {
	// 	a, err := res.Recv();
	// 	if err != nil {
	// 		panic(err);
	// 	}
	// 	fmt.Println(a);
	// }

	// puts, err := c.PutStream(context.Background());
	// if err != nil {
	// 	fmt.Println(err);
	// }

	// i := 0;

	// for {
	// 	i++;
	// 	puts.Send(&streampb.StreamReq{
	// 		Data: fmt.Sprintf("sdsd %d", i),
	// 	});
	// 	time.Sleep(time.Second);
	// 	if i>10 {
	// 		fmt.Println(err);
	// 		break;
	// 	}
	// }
//  allStr, _ := c.AllStream(ctx);
// 	wg := sync.WaitGroup{};
// 	wg.Add(2);
// 	go func() {
// 		defer wg.Done();
// 		for {
// 			data, err := allStr.Recv();
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(data.Data);
// 		}
// 	}();

// 	go func() {
// 		defer wg.Done()
// 		for {
// 			_ = allStr.Send(&streampb.StreamReq{Data: "我是客户端的数据"});
// 			time.Sleep(time.Second);
// 		}
// 	}();
// 		wg.Wait()
}