package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) GetAllProductsStorage(ctx context.Context) ([]*models.Product, error) {
	var res []*models.Product

	builder := sq.Select("name", "price").
		From("products").
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, description string
		var price float64
		var count uint32

		if err = rows.Scan(&name, &price); err != nil {
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
