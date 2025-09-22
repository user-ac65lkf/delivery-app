package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) UpdateBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) (int64, error) {
	q := sq.Update("basket").
		Set("amount", req.Count).
		Where(sq.Eq{"user_id": req.UserId, "product_id": req.ProductId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	res, err := q.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
