package service

import (
	"context"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"github.com/Shemistan/uzum_delivery/internal/storage"
	pbLogin "github.com/Shemistan/uzum_delivery/pkg/login_v1"
)

type IService interface {
	VerifyToken(ctx context.Context) (int, error)
	GetOrder(ctx context.Context, id int64, courierId int64) (*models.Order, error)
	GetOrders(ctx context.Context, coord *models.Coordinate) ([]*models.IdDist, error)
	CloseOrder(ctx context.Context, id int64) error
}

func NewService(storage storage.IStorage, loginClient pbLogin.LoginV1Client) IService {
	return &service{
		Storage:     storage,
		loginClient: loginClient,
	}
}

type service struct {
	loginClient pbLogin.LoginV1Client
	Storage     storage.IStorage
}
