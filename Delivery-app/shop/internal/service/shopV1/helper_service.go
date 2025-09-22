package shopV1

import (
	"context"
	loginPb "github.com/Shemistan/uzum_shop/pkg/loginV1"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func (s *shopSystemService) GetUserIdFromLoginServ(ctx context.Context) (int, error) {
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
