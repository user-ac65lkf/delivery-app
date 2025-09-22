package postgres

import (
	"context"
	"database/sql"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (r *repo) GetOrder(ctx context.Context, id int64, courierId int64) (*models.Order, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("Unable to begin transaction because %s", err)
	}
	defer tx.Rollback()

	timeNow := time.Now().Format("2006-01-02 15:04:05")

	sqlStm := `UPDATE "orders" SET start_at=$1, courier_id=$2, delivery_status=$3 WHERE id=$4`

	_, err = tx.ExecContext(ctx, sqlStm, timeNow, courierId, "InProgress", id)
	if err != nil {
		log.Fatalf("not able to update orders table because %s", err)
	}

	var order models.Order

	row := tx.QueryRowContext(ctx, `SELECT orders.id, orders.address, orders.coordinate_address_x,
    orders.coordinate_address_y, orders.coordinates_point_x, orders.coordinates_point_y, orders.start_at, users.name, users.phone 
											FROM orders
                                 			JOIN users 
										 	ON orders.user_id = users.id
										 	WHERE orders.id = $1;`, id)
	err = row.Scan(&order.ID, &order.Address, &order.CoordinateAddress.Latitude, &order.CoordinateAddress.Longitude,
		&order.CoordinatePickup.Latitude, &order.CoordinatePickup.Longitude, &order.StartedAt, &order.Name, &order.Phone)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Printf("Unable to retrieve with ID- %v", id)
		return nil, err
	case err != nil:
		log.Fatalf(err.Error())
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &order, nil
}
