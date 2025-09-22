package service

import (
	"context"
	loginPb "github.com/Shemistan/uzum_delivery/pkg/login_v1"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func (s *service) VerifyToken(ctx context.Context) (int, error) {
	emp := &loginPb.Check_Request{EndpointAddress: ""}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	check, err := s.loginClient.Check(ctx, emp)
	if err != nil {
		return -1, err
	}

	userId, err := strconv.Atoi(check.UserId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
