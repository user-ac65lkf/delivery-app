package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) DeleteProduct(ctx context.Context, id int) error {
	q := sq.Delete("products").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
