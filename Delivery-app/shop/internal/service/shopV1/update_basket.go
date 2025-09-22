package shopV1

import (
	"context"
	"fmt"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) UpdateBasketService(ctx context.Context, req *models.AddProductToBasketModel) error {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return err
	}

	req.UserId = userId

	res, err := s.storage.UpdateBasketStorage(ctx, req)
	if err != nil {
		return err
	}

	if res == 0 {
		return fmt.Errorf("%v rows affected", res)
	}

	return nil
}
