package shopV1

import (
	"context"
	shop_system "github.com/Shemistan/uzum_shop/internal/service/shopV1"
	"github.com/Shemistan/uzum_shop/pkg/shopV1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Shop struct {
	shopV1.UnimplementedShopServer
	ShopService shop_system.IShopSystemService
}

func (a *Shop) Healthz(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}
