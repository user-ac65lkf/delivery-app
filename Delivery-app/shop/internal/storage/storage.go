package storage

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
)

type IStorage interface {
	GetProductStorage(ctx context.Context, prodId uint32) (*models.Product, error)
	GetAllProductsStorage(ctx context.Context) ([]*models.Product, error)
	AddProductToBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) error
	UpdateBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) (int64, error)
	DeleteBasketStorage(ctx context.Context, req *models.DeleteFomBasked) error
	GetBasketStorage(ctx context.Context, userId int) ([]*models.BasketItem, error)
	CreateOrderStorage(ctx context.Context, req *models.Order) (uint32, error)
	CancelOrderStorage(ctx context.Context, orderId uint32) error
	GetProductCountStorage(ctx context.Context, prodId uint32) (int, error)
	CalculateProductCountStorage(ctx context.Context, prodId int, count int) (int64, error)
	GetAddress(ctx context.Context, userId int) (string, error)
	GetItemsFromBasket(ctx context.Context, userId int) ([]*models.GetFromBasket, error)
	CreateOrderDetails(ctx context.Context, order_id int, items []*models.GetFromBasket) error
	CancelOrderDetailsStorage(ctx context.Context, order_id int) ([]*models.GetFromBasket, error)
}
