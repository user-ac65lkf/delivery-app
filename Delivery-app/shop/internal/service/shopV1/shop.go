package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	repo "github.com/Shemistan/uzum_shop/internal/storage"
	loginPb "github.com/Shemistan/uzum_shop/pkg/loginV1"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IShopSystemService interface {
	GetUserIdFromLoginServ(ctx context.Context) (int, error)
	GetProductService(ctx context.Context, request *pb.GetProduct_Request) (*models.Product, error)
	GetAllProductsService(ctx context.Context) ([]*models.Product, error)
	AddProductToBasketService(ctx context.Context, req *models.AddProductToBasketModel) error
	UpdateBasketService(ctx context.Context, req *models.AddProductToBasketModel) error
	DeleteBasketService(ctx context.Context, req *models.DeleteFomBasked) error
	GetBasketService(ctx context.Context) ([]*models.BasketItem, error)
	CreateOrderService(ctx context.Context, req *models.Order) (int32, error)
	CancelOrderService(ctx context.Context, orderId uint32) error
}

type shopSystemService struct {
	storage     repo.IStorage
	loginClient loginPb.LoginV1Client
	kProducer   *kafka.Producer
}

func NewShopSystemService(storage repo.IStorage, loginClient loginPb.LoginV1Client, kProducer *kafka.Producer) IShopSystemService {
	return &shopSystemService{
		storage:     storage,
		loginClient: loginClient,
		kProducer:   kProducer,
	}
}
