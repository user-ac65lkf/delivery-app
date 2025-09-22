package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_delivery/internal/models"
)

func (r *repo) GetOrders(ctx context.Context) ([]*models.GetOrders, error) {

	var orders []*models.GetOrders
	builder := sq.Select("id", "coordinate_address_x", "coordinate_address_y",
		"coordinates_point_x", "coordinates_point_y").
		From("orders").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var dropX, dropY, pickX, pickY float64

		if err = rows.Scan(&id, &dropX, &dropY, &pickX, &pickY); err != nil {
			return nil, err
		}

		orders = append(orders, &models.GetOrders{
			ID: id,
			DropCoord: &models.Coordinate{
				Latitude:  dropX,
				Longitude: dropY,
			},
			PickUpCoord: &models.Coordinate{
				Latitude:  pickX,
				Longitude: pickY,
			},
		})
	}

	return orders, nil
}
