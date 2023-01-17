package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/client/profile"
	"coolcar/rental/trip/dao"
	"coolcar/shared/server"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger,err := server.NewZapLogger();
	
	if err != nil {
		log.Fatalf("cannot find logger %v", err)
	}
	c := context.Background()
	MongoClient, err := mongo.Connect(c,options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false&directConnection=true"))

	if err != nil {
		logger.Fatal("cannot find mongodb", zap.Error(err))
	}
	
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name: "rental",
		Addr: ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{},
				ProfileManager: &profile.Manager{},
				POIManager: &poi.Manager{},
				Mongo: dao.NewMongo(MongoClient.Database("coolcar")),
				Logger: logger,
			});
		},

	})
	// lis, err := net.Listen("tcp",":8082");
	if err != nil {
		logger.Sugar().Fatal(err)
		logger.Fatal("cannot start server", zap.Error(err))
	}
	// c := context.Background()
	// MongoClient, err := mongo.Connect(c,options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false&directConnection=true"))

	// if err != nil {
	// 	logger.Fatal("cannot find mongodb", zap.Error(err))
	// }
	
  // pkFile,err := os.OpenFile("private.key", os.O_RDWR,os.ModeAppend);

	// if err != nil {
	// 	panic(err)
	// }
	// pkBytes, err := ioutil.ReadAll(pkFile)
	// if err != nil {
	// 	panic(err)
	// }
	// privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	// if err != nil {
	// 	panic(err)
	// }
	// in, err := auth.Interceptor("shared/auth/public.key");
	// if err != nil {
	// 	logger.Fatal("cannot create auth interceptor")
	// }
	// s :=grpc.NewServer(grpc.UnaryInterceptor(in));
	// rentalpb.RegisterTripServiceServer(s, &trip.Service{
	// 	Logger: logger,
	// });
	// s.Serve(lis)
	// authpb.RegisterAuthServiceServer(s, &auth.Service{
	// 	OpenIDResolver: &wechat.Service{
	// 		AppID: "wxb85f823075100a64",
	// 		AppSecret: "4e1fa08a4b270099497f4935e57b916d",
	// 	},
	// 	Logger: logger,
	// 	Mongo: dao.NewMongo(MongoClient.Database("coolcar")),
	// 	TokenExpire: 2 * time.Hour,
	// 	TokenGenerator: token.NewJWTTokenGen("coolcar/auth",privatekey) ,
	// })

	
}

// func newZapLogger() (*zap.Logger, error) {
// 	cfg := zap.NewDevelopmentConfig();
// 	cfg.EncoderConfig.TimeKey = "";

// 	return cfg.Build()

// }