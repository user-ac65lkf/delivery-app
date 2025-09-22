package postgres

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (r *repo) CloseOrder(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("Unable to begin transaction because %s", err)
	}
	defer tx.Rollback()

	var status string

	row := tx.QueryRowContext(ctx, `SELECT orders.delivery_status
    										FROM orders
    										    WHERE orders.id=$1;`, id)

	err = row.Scan(&status)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Printf("Unable to retrieve with ID- %v", id)
		return err
	case err != nil:
		log.Fatalf(err.Error())
		return err
	}

	if status != "InProgress" {
		return errors.New("order status not InProgress")
	}

	timeNow := time.Now().Format("2006-01-02 15:04:05")

	sqlStm := `UPDATE "orders" SET delivery_at=$1, delivery_status=$2 WHERE id=$3`

	_, err = tx.ExecContext(ctx, sqlStm, timeNow, "Close", id)
	if err != nil {
		log.Fatalf("not able to update orders table because %s", err)
		return err
	}

	sqlDelStm := `DELETE FROM "orders_details"  WHERE order_id = $1`

	_, err = tx.ExecContext(ctx, sqlDelStm, id)
	if err != nil {
		log.Fatalf("unable to delete orders_details table row because %s", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
