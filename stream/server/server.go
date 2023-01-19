package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	streampb "coolcar/stream/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const PORT = ":50052";

type server struct {}

func (s *server)NormalSteam(ctx context.Context, req *streampb.StreamReq ) (*streampb.StramRes, error) {

	return nil, nil

}

// 服务端流模式
func (s *server)GetStream(req *streampb.StreamReq, res streampb.Greeter_GetStreamServer) error {
	i :=0;

	for {
		i++;
		_ = res.Send(&streampb.StramRes{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		});

		time.Sleep(time.Second);

		if i > 10 {
			break;
		}
		
	}

	return nil;
}
// 服务端流模式
func (s *server)PutStream(cliStr streampb.Greeter_PutStreamServer) (error) {
	for {
		a, err := cliStr.Recv();
		if err != nil {
			fmt.Println(err);
			break;
		}
		fmt.Println(a.Data);

	}
	return nil;
}
// 双流模式
func (s *server)AllStream(allStr streampb.Greeter_AllStreamServer) ( error) {

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
			_ = allStr.Send(&streampb.StramRes{Data: "我是服务器的数据"});
			time.Sleep(time.Second);
		}
	}();
		wg.Wait()
	return nil;
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("cannot listen", zap.Error(err))
	}

	var opts []grpc.ServerOption
	// if c.AuthPublicKeyFile != "" {
	// 	in, err := auth.Interceptor(c.AuthPublicKeyFile)
	// 	if err != nil {
	// 		c.Logger.Fatal("cannot create auth interceptor", nameField, zap.Error(err))
	// 	}
	// 	opts = append(opts, grpc.UnaryInterceptor(in))
	// }

	s := grpc.NewServer(opts...)
	streampb.RegisterGreeterServer(s, &server{})
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	// c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))

	err = s.Serve(lis);

	if err != nil {
		panic(err)
	}

}
