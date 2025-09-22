package app

import (
	"context"
	shop_system_v1 "github.com/Shemistan/uzum_shop/internal/api/shopV1"

	pbgrpc "github.com/Shemistan/uzum_shop/pkg/shopV1"
	gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func (a *App) initDB() {
	sqlConnectionString := a.getSqlConnectionString()
	var err error
	a.db, err = sqlx.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal("failed to opening connection to db: ", err.Error())
	}

	// Проверка соединения с базой данных
	if err = a.db.Ping(); err != nil {
		log.Fatal("failed to connect to the database: ", err.Error())
	}
}

func (a *App) initGRPCServer() {
	a.grpcShopServer = grpc.NewServer()
	pbgrpc.RegisterShopServer(
		a.grpcShopServer,
		&shop_system_v1.Shop{
			ShopService: a.getShopSystemService(),
		},
	)
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.muxAuth = gateway_runtime.NewServeMux()
	optsShop := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pbgrpc.RegisterShopHandlerFromEndpoint(ctx, a.muxAuth, a.appConfig.App.PortAuthGRPC, optsShop)
	if err != nil {
		return err
	}
	return nil
}
