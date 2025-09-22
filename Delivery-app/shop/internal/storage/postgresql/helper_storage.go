package postgresql

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) GetProductCountStorage(ctx context.Context, prodId uint32) (int, error) {
	var count int
	q := sq.Select("count").
		From("products").
		Where(sq.Eq{"id": int(prodId)}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	err := q.QueryRowContext(ctx).
		Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (s *storage) CalculateProductCountStorage(ctx context.Context, prodId int, count int) (int64, error) {
	q := sq.Update("products").
		Set("count", count).
		Where(sq.Eq{"id": prodId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	res, err := q.ExecContext(ctx)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (s *storage) GetAddress(ctx context.Context, userId int) (string, error) {
	var address sql.NullString
	q := sq.Select("address").
		From("users").
		Where(sq.Eq{"id": userId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	err := q.QueryRowContext(ctx).
		Scan(&address)

	if err != nil {
		return "", err
	}

	return address.String, nil
}

func (s *storage) GetItemsFromBasket(ctx context.Context, userId int) ([]*models.GetFromBasket, error) {
	var res []*models.GetFromBasket
	q := sq.Select("product_id", "amount").
		From("basket").
		Where(sq.Eq{"user_id": userId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := q.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		var prodId, amount int

		if err = rows.Scan(&prodId, &amount); err != nil {
			return nil, err
		}
		res = append(res, &models.GetFromBasket{
			ProductId: prodId,
			Count:     amount,
		})
		counter++
	}

	if counter == 0 {
		return nil, errors.New("basket is empty")
	}

	return res, nil
}

func (s *storage) CreateOrderDetails(ctx context.Context, order_id int, items []*models.GetFromBasket) error {
	q := sq.Insert("orders_details").
		Columns("order_id", "product_id", "quantity").
		RunWith(s.db).PlaceholderFormat(sq.Dollar)

	for _, item := range items {
		q = q.Values(order_id, item.ProductId, item.Count)
	}

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) CancelOrderDetailsStorage(ctx context.Context, order_id int) ([]*models.GetFromBasket, error) {

	var res []*models.GetFromBasket
	qx := sq.Select("product_id", "quantity").
		From("orders_details").
		Where(sq.Eq{"order_id": order_id}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := qx.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var prodId, quantity int

		if err = rows.Scan(&prodId, &quantity); err != nil {
			return nil, err
		}
		res = append(res, &models.GetFromBasket{
			ProductId: prodId,
			Count:     quantity,
		})
	}

	q := sq.Delete("orders_details").
		Where(sq.Eq{"order_id": order_id}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err = q.ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
