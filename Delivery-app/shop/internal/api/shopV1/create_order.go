package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"time"
)

func (serv *Shop) CreateOrder(ctx context.Context, req *pb.Order_Request) (*pb.Order_Response, error) {
	orderModel := &models.Order{
		Address:              req.Address,
		Coordinate_address_x: req.DropX,
		Coordinate_address_y: req.DropY,
		Coordinates_point_x:  req.TakeX,
		Coordinates_point_y:  req.TakeY,
		Create_at:            time.Now().Format("2006-01-02 15:04:05"),
		Delivery_status:      "Awaiting Shipment",
	}

	resp, err := serv.ShopService.CreateOrderService(ctx, orderModel)
	if err != nil {
		return &pb.Order_Response{}, err
	}

	return &pb.Order_Response{
		OrderId: uint32(resp),
	}, nil
}
