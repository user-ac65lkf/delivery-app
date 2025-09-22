package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *service) GetProduct(ctx context.Context, productId int) (*models.Product, error) {
	return s.storage.GetProduct(ctx, productId)
}
