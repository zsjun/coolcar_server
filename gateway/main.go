package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	log.SetFlags(log.Lshortfile)
	startGRPCGateway()
	// // 监听端口
	// lis, err := net.Listen("tcp", ":8081")
	// if(err != nil) {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// // 创建grpc服务
	// grpcServer := grpc.NewServer()
	
	// authpb.RegisterAuthServiceServer(grpcServer, &auth.Service{})
	
	// // if err != nil {
	// // 	fmt.Println("服务启动失败", err)
	// // 	return
	// // }
	// log.Fatal(grpcServer.Serve(lis))
	// fmt.Println("连接成功")

}

func startGRPCGateway() {
	lg, err := server.NewZapLogger();
	if err != nil {
		log.Fatalf("cannot create zap:logger: %v", err)
	}
	// 建立空的上下文
	c := context.Background();
	c, cancel := context.WithCancel(c);
	defer cancel()
	// 分发器，当有服务的时候，分发到哪个服务器上面
	// customMar := jsonpb.Marshaler{
	// 	EnumsAsInts: true,
	// }
	// mar := runtime.JSONPb(customMar)
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:true,
				UseProtoNames: true,
			},
		},
	))
	serverConfig := []struct{
		name string;
		addr string;
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:"auth",
			addr:"localhost:8081",
			registerFunc:authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:"trip",
			addr:"localhost:8082",
			registerFunc:rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:"profile",
			addr:"localhost:8082",
			registerFunc:rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},

	}
	for _, s := range serverConfig {
		err := s.registerFunc(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatalf("cannot register service %s: %v",s.name,err)
		}
	}
	// err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
	// if(err != nil) {
	// 	log.Fatalf("cannot start grpc gateway %v", err)
	// }
	// err = rentalpb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8082", []grpc.DialOption{grpc.WithInsecure()})
	// if(err != nil) {
	// 	log.Fatalf("cannot start grpc gateway 8082 %v", err)
	// }
	// 8080的请求分发到8081的服务器上面
	addr := ":8080";
	lg.Sugar().Infof("grpc server started ad: %s", addr);
	err = http.ListenAndServe(addr, mux);
	if err != nil {
		log.Fatalf("cannot listen to 8080 %v", err)
	}
}