package storage

import (
	"context"
	"github.com/Shemistan/uzum_delivery/internal/models"
)

type IStorage interface {
	GetOrder(ctx context.Context, id int64, courierId int64) (*models.Order, error)
	GetOrders(ctx context.Context) ([]*models.GetOrders, error)
	CloseOrder(ctx context.Context, id int64) error
}
