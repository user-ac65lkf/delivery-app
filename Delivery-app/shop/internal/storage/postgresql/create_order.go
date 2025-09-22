package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
	"github.com/lib/pq"
)

func (s *storage) CreateOrderStorage(ctx context.Context, req *models.Order) (uint32, error) {
	var orderId int
	q := sq.Insert("orders").
		Columns("user_id", "products_id", "address", "coordinate_address_x", "coordinate_address_y",
			"coordinates_point_x", "coordinates_point_y", "create_at", "delivery_status").
		Values(req.User_id, pq.Array(req.Products_id), req.Address, req.Coordinate_address_x, req.Coordinate_address_y,
			req.Coordinates_point_x, req.Coordinates_point_y, req.Create_at, req.Delivery_status).
		Suffix("RETURNING \"id\"").
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)
	err := q.QueryRow().Scan(&orderId)
	if err != nil {
		return 0, err
	}

	return uint32(orderId), nil
}
