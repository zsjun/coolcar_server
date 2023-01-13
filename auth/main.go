package main

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger,err := zap.NewDevelopment();
	newZapLogger()
	if(err != nil) {
		log.Fatalf("cannot find logger %v", err)
	}
	lis, err := net.Listen("tcp",":8081");
	if(err != nil) {
		logger.Fatal("cannot err", zap.Error(err))
	}
	s :=grpc.NewServer();

	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
	})

	s.Serve(lis)
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig();
	cfg.EncoderConfig.TimeKey = "";

	return cfg.Build()


}