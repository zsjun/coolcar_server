package main

import (
	"context"
	streampb "coolcar/stream/api/gen/v1"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052",grpc.WithInsecure());

	if err != nil {
		panic(err);
	}
	defer conn.Close();

	c := streampb.NewGreeterClient(conn);

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
 allStr, _ := c.AllStream(context.Background());
	wg := sync.WaitGroup{};
	wg.Add(2);
	go func() {
		defer wg.Done();
		for {
			data, err := allStr.Recv();
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data.Data);
		}
	}();

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&streampb.StreamReq{Data: "我是客户端的数据"});
			time.Sleep(time.Second);
		}
	}();
		wg.Wait()
}