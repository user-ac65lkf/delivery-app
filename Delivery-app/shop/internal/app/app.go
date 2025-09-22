package app

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"runtime"

	"github.com/Shemistan/uzum_shop/dev"
	"github.com/Shemistan/uzum_shop/internal/models"
	shop_system "github.com/Shemistan/uzum_shop/internal/service/shopV1"
	"github.com/Shemistan/uzum_shop/internal/storage/postgresql"
	loginPb "github.com/Shemistan/uzum_shop/pkg/loginV1"
	gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type App struct {
	appConfig *models.Config
	muxAuth   *gateway_runtime.ServeMux

	grpcShopServer    *grpc.Server
	shopSystemService shop_system.IShopSystemService

	db *sqlx.DB
}

func NewApp(ctx context.Context) (*App, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	a := &App{}
	a.setConfig()
	a.initDB()
	a.initGRPCServer()
	if err := a.initHTTPServer(ctx); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) setConfig() {
	if dev.DEBUG {
		err := dev.SetConfig()
		if err != nil {
			log.Fatal("failed to get config:", err.Error())
		}
	}
	conf := models.Config{}

	envconfig.MustProcess("", &conf)

	a.appConfig = &conf
}

func (a *App) getShopSystemService() shop_system.IShopSystemService {
	storage := postgresql.NewStorage(a.db)

	conn, err := grpc.Dial("auth:9081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	//defer conn.Close()

	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092", // адресс
	}
	// Создаем продюсера
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Println("failed to create producer")
	}
	producer.Flush(15 * 1000)

	loginCl := loginPb.NewLoginV1Client(conn)

	if a.shopSystemService == nil {
		a.shopSystemService = shop_system.NewShopSystemService(storage, loginCl, producer)
	}

	return a.shopSystemService
}

func (a *App) getSqlConnectionString() string {
	sqlConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%v",
		a.appConfig.DB.User,
		a.appConfig.DB.Password,
		a.appConfig.DB.Host,
		a.appConfig.DB.Port,
		a.appConfig.DB.Database,
		a.appConfig.DB.SSLMode,
	)
	return sqlConnectionString
}
