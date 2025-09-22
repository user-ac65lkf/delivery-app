package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) GetAllProductsService(ctx context.Context) ([]*models.Product, error) {
	response, err := s.storage.GetAllProductsStorage(ctx)
	if err != nil {
		return nil, err
	}

	return response, nil
}
