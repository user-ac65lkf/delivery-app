package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) GetBasketService(ctx context.Context) ([]*models.BasketItem, error) {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return nil, err
	}

	response, err := s.storage.GetBasketStorage(ctx, userId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
