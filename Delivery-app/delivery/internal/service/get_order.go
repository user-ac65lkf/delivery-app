package service

import (
	"context"
	"github.com/Shemistan/uzum_delivery/internal/models"
)

func (s *service) GetOrder(ctx context.Context, id int64, courierId int64) (*models.Order, error) {

	return s.Storage.GetOrder(ctx, id, courierId)
}
