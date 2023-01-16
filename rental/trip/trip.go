package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct{
	Logger *zap.Logger;
}


func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error){
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		fmt.Printf("sdds %v\n",err);
		return nil, err;
	}
	s.Logger.Info("create trip", zap.String("start",req.Start), zap.String("accound_id", string(aid)));

	return nil, status.Error(codes.Unimplemented, "", )
}
