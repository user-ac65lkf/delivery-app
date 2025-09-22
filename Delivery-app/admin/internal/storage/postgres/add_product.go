package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (r *repo) AddProduct(ctx context.Context, req *models.Product) (int64, error) {
	builder := sq.Insert("products").
		Columns("name", "description", "price", "count").
		Values(req.Name, req.Description, req.Price, req.Count).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar).
		Suffix(`RETURNING 
				"id"
				`)

	var id int64

	err := builder.QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
