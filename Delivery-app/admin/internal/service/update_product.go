package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *service) UpdateProduct(ctx context.Context, req *models.Product) error {
	return s.storage.UpdateProduct(ctx, req)
}
