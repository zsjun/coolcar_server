package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/wechat"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	c := context.Background()
	MongoClient, err := mongo.Connect(c,options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false&directConnection=true"))

	if err != nil {
		logger.Fatal("cannot find mongodb", zap.Error(err))
	}

	s :=grpc.NewServer();

	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID: "wxb85f823075100a64",
			AppSecret: "4e1fa08a4b270099497f4935e57b916d",
		},
		Logger: logger,
		Mongo: dao.NewMongo(MongoClient.Database("coolcar")),
	})

	s.Serve(lis)
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig();
	cfg.EncoderConfig.TimeKey = "";

	return cfg.Build()


}