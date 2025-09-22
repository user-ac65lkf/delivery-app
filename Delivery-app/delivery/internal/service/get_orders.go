package service

import (
	"context"
	"fmt"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"slices"
	"sort"
)

func (s *service) GetOrders(ctx context.Context, coord *models.Coordinate) ([]*models.IdDist, error) {
	ordersMap := make(map[int64]float64)

	var ordersSlice []*models.IdDist
	orders, err := s.Storage.GetOrders(ctx)

	for _, val := range orders {
		fmt.Printf("%v, ", val.ID)
	}

	var total float64 = 0
	values := make([]float64, 0, len(orders))
	checkedId := make([]int64, 0, len(orders))

	for _, v := range orders {
		distPickAndDrop := Distance(v.DropCoord.Latitude, v.DropCoord.Longitude, v.PickUpCoord.Latitude, v.PickUpCoord.Longitude)

		distPickAndCourier := Distance(coord.Latitude, coord.Longitude, v.PickUpCoord.Latitude, v.PickUpCoord.Longitude)

		total = distPickAndDrop + distPickAndCourier
		ordersMap[v.ID] = total
		values = append(values, total)

	}

	sort.Float64s(values)

	var count int

	for _, dist := range values {

		for k, distance := range ordersMap {

			if slices.Contains(checkedId, k) {
				continue
			}

			fmt.Printf("\nfor k, val := range ordersMap { KK- %v \n", k)
			if dist == distance {
				checkedId = append(checkedId, k)
				ordersSlice = append(ordersSlice, &models.IdDist{
					Id:   k,
					Dist: dist,
				})
			}
		}

		count++
		if count == 5 {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return ordersSlice, nil
}
