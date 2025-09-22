package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/storage"
	desc "github.com/Shemistan/uzum_admin/pkg/login_v1"
)

type IService interface {
	VerifyToken(ctx context.Context) (int, error)
	AddProduct(ctx context.Context, req *models.Product) (int64, error)
	UpdateProduct(ctx context.Context, req *models.Product) error
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	GetProduct(ctx context.Context, productId int) (*models.Product, error)
	DeleteProduct(ctx context.Context, productId int) error
}

func NewService(storage storage.IStorage, loginClient desc.LoginV1Client) IService {
	return &service{
		storage:     storage,
		loginClient: loginClient,
	}
}

type service struct {
	storage     storage.IStorage
	loginClient desc.LoginV1Client
}
