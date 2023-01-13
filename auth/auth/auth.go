package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
)
type Service struct {
	Logger zap.Logger
	// authpb.UnimplementedTripServiceServer
}


func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("receive code",zap.String("code", req.Code));
	return nil,nil;
}