package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (r *repo) GetProduct(ctx context.Context, productId int) (*models.Product, error) {
	var product models.Product
	q := sq.Select("id", "name", "description", "price",
		"count").
		From("products").
		Where(sq.Eq{"id": productId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := q.QueryRowContext(ctx).
		Scan(&product.ID, &product.Name, &product.Description,
			&product.Price, &product.Count)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
