package storage

import (
	"context"

	"github.com/Shemistan/uzum_admin/internal/models"
)

type IStorage interface {
	AddProduct(ctx context.Context, req *models.Product) (int64, error)
	UpdateProduct(ctx context.Context, req *models.Product) error
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	GetProduct(ctx context.Context, id int) (*models.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}
