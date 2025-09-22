package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *service) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	return s.storage.GetAllProducts(ctx)
}
