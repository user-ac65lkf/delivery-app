package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *storage) CancelOrderStorage(ctx context.Context, orderId uint32) error {
	q := sq.Delete("orders").
		Where(sq.Eq{"id": orderId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
