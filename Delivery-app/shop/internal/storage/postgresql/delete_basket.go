package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *storage) DeleteBasketStorage(ctx context.Context, req *models.DeleteFomBasked) error {
	q := sq.Delete("basket").
		Where(sq.Eq{"product_id": req.ProductId, "user_id": req.UserId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
