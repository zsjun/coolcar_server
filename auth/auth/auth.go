package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"log"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Service struct {
	OpenIDResolver OpenIDResolver;
	Logger *zap.Logger;
	Mongo *dao.Mongo;
	TokenGenerator TokenGenerator;
	TokenExpire    time.Duration;
}

type OpenIDResolver interface {
	Resolve(code string) (string,error);
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration)(string, error)
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
	tkn, err := s.TokenGenerator.GenerateToken(string(accountID),s.TokenExpire);

	if err != nil {
		s.Logger.Error("cannot generate token")
		return nil, status.Error(codes.Internal,"")
	}
	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn: int32(s.TokenExpire.Seconds()),
	},nil;
}