package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (r *repo) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	var res []*models.Product

	builder := sq.Select("id", "name", "description", "price", "count").
		From("products").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var price float64
		var count uint64
		var name, description string

		if err = rows.Scan(&id, &name, &description, &price, &count); err != nil {
			return nil, err
		}

		res = append(res, &models.Product{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Count:       count,
		})
	}

	return res, nil
}
