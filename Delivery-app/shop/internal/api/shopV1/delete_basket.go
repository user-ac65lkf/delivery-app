package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) DeleteBasket(ctx context.Context, req *pb.DeleteBasket_Request) (*empty.Empty, error) {
	deleteReq := &models.DeleteFomBasked{
		ProductId: req.ProductId,
	}

	err := serv.ShopService.DeleteBasketService(ctx, deleteReq)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
