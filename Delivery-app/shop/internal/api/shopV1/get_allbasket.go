package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) GetBasket(ctx context.Context, _ *empty.Empty) (*pb.GetBasket_Response, error) {
	response, err := serv.ShopService.GetBasketService(ctx)
	if err != nil {
		return nil, err
	}

	var res []*pb.BasketItem

	for _, v := range response {

		res = append(res, &pb.BasketItem{
			Id:        v.Id,
			UserId:    v.UserId,
			ProductId: v.ProductId,
			Count:     v.Count,
		})
	}

	respond := &pb.GetBasket_Response{
		AllBasket: res,
	}

	return respond, nil
}
