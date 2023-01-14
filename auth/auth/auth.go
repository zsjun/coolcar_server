package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Service struct {
	OpenIDResolver OpenIDResolver;
	Logger *zap.Logger;
	Mongo *dao.Mongo;
}

type OpenIDResolver interface {
	Resolve(code string) (string,error);
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("receive code",zap.String("code", req.Code));
	openID,err := s.OpenIDResolver.Resolve(req.Code)
	if err !=nil  {
		log.Fatalf("cannot get service %v", err)
		return nil, status.Errorf(codes.Unavailable, "cannot resolce openId %v", err)
	}
	
	accountID, err := s.Mongo.ResolveAccountID(c, openID);
	if err != nil {
		return nil, status.Error(codes.Internal,"")
	}
	return &authpb.LoginResponse{
		AccessToken: string(accountID),
		ExpiresIn: 7200,
	},nil;
}