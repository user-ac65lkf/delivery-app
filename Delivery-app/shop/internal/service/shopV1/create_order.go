package shopV1

import (
	"context"
	"errors"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) CreateOrderService(ctx context.Context, req *models.Order) (int32, error) {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return 0, err
	}

	req.User_id = userId

	getItemsFromBasket, err := s.storage.GetItemsFromBasket(ctx, userId)
	if err != nil {
		return 0, err
	}

	stockMap := make(map[int]int)

	for _, v := range getItemsFromBasket {
		stockCounts, err := s.storage.GetProductCountStorage(ctx, uint32(v.ProductId))
		if err != nil {
			return 0, err
		}

		if stockCounts < v.Count {
			return 0, errors.New("not enough stock")
		}
		stockMap[v.ProductId] = stockCounts
	}

	var address string = req.Address

	if req.Address == "" {
		address, err = s.storage.GetAddress(ctx, userId)
		if err != nil {
			return 0, err
		}

		if address == "" {
			return 0, errors.New("no address provided")
		}
	}

	respOrderId, err := s.storage.CreateOrderStorage(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.storage.CreateOrderDetails(ctx, int(respOrderId), getItemsFromBasket)
	if err != nil {
		return 0, err
	}

	for _, item := range getItemsFromBasket {

		deleteReq := &models.DeleteFomBasked{
			UserId:    userId,
			ProductId: uint32(item.ProductId),
		}

		err = s.storage.DeleteBasketStorage(ctx, deleteReq)
		if err != nil {
			return 0, err
		}
	}

	return int32(respOrderId), nil
}
