package shopV1

import (
	"context"
)

func (s *shopSystemService) CancelOrderService(ctx context.Context, orderId uint32) error {
	resp, err := s.storage.CancelOrderDetailsStorage(ctx, int(orderId))
	if err != nil {
		return err
	}

	err = s.storage.CancelOrderStorage(ctx, orderId)
	if err != nil {
		return err
	}

	for _, v := range resp {
		stockCounts, err := s.storage.GetProductCountStorage(ctx, uint32(v.ProductId))
		if err != nil {
			return err
		}

		countToSet := stockCounts + v.Count

		_, err = s.storage.CalculateProductCountStorage(ctx, v.ProductId, countToSet)
		if err != nil {
			return err
		}
	}

	return nil
}
