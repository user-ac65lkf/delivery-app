package service

import (
	"context"
)

func (s *service) CloseOrder(ctx context.Context, id int64) error {

	return s.Storage.CloseOrder(ctx, id)
}
