package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) GetProductService(ctx context.Context, req *pb.GetProduct_Request) (*models.Product, error) {
	response, err := s.storage.GetProductStorage(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
