package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger,err := server.NewZapLogger();
	// newZapLogger()
	if err != nil {
		log.Fatalf("cannot find logger %v", err)
	}

	c := context.Background()
	MongoClient, err := mongo.Connect(c,options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false&directConnection=true"))

	if err != nil {
		logger.Fatal("cannot find mongodb", zap.Error(err))
	}
	
  pkFile,err := os.OpenFile("./auth/private.key", os.O_RDWR,os.ModeAppend);

	if err != nil {
		panic(err)
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		panic(err)
	}
	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		panic(err)
	}

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:"auth",
		Addr: ":8081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIDResolver: &wechat.Service{
					AppID: "wxb85f823075100a64",
					AppSecret: "4e1fa08a4b270099497f4935e57b916d",
				},
				Logger: logger,
				Mongo: dao.NewMongo(MongoClient.Database("coolcar")),
				TokenExpire: 2 * time.Hour,
				TokenGenerator: token.NewJWTTokenGen("coolcar/auth",privatekey) ,
			})
		},
	})
	if err !=nil {
		logger.Sugar().Fatal(err)
	}
}
