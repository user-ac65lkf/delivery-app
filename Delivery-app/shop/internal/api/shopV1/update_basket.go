package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) UpdateBasket(ctx context.Context, req *pb.UpdateBasket_Request) (*empty.Empty, error) {
	updateBask := &models.AddProductToBasketModel{
		ProductId: req.ProductId,
		Count:     req.Count,
	}

	err := serv.ShopService.UpdateBasketService(ctx, updateBask)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
