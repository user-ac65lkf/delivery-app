package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) CancelOrder(ctx context.Context, req *pb.CancelOrder_Request) (*empty.Empty, error) {
	err := serv.ShopService.CancelOrderService(ctx, req.OrderId)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
