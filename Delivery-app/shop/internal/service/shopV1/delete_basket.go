package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) DeleteBasketService(ctx context.Context, req *models.DeleteFomBasked) error {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return err
	}

	req.UserId = userId

	err = s.storage.DeleteBasketStorage(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
