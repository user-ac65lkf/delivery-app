package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) AddProductToBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) error {

	q := sq.Insert("basket").
		Columns("user_id", "product_id", "amount").
		Values(req.UserId, req.ProductId, req.Count).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
