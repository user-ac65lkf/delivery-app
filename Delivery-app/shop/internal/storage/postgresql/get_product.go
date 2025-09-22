package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) GetProductStorage(ctx context.Context, prodId uint32) (*models.Product, error) {
	var product models.Product
	q := sq.Select("id", "name", "description", "price", "count").
		From("products").
		Where(sq.Eq{"id": int(prodId)}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	err := q.QueryRowContext(ctx).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
