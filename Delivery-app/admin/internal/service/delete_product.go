package service

import (
	"context"
)

func (s *service) DeleteProduct(ctx context.Context, productId int) error {
	return s.storage.DeleteProduct(ctx, productId)
}
