package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (serv *Shop) GetProduct(ctx context.Context, req *pb.GetProduct_Request) (*pb.GetProduct_Response, error) {
	response, err := serv.ShopService.GetProductService(ctx, req)
	if err != nil {
		return &pb.GetProduct_Response{}, err
	}

	return &pb.GetProduct_Response{
		Product: &pb.Product{
			Id:          uint32(response.ID),
			Name:        response.Name,
			Description: response.Description,
			Price:       response.Price,
			Count:       response.Count,
		},
	}, nil


}
