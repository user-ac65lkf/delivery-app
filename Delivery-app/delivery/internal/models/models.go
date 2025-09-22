package models

import "time"

type Order struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	CoordinateAddress Coordinate `json:"coordinateAddress"`
	CoordinatePickup  Coordinate `json:"coordinatePickup"`

	Meta      string    `json:"meta,omitempty"`
	StartedAt time.Time `json:"startedAt"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetOrders struct {
	ID          int64       `json:"id"`
	DropCoord   *Coordinate `json:"dropCoordinate"`
	PickUpCoord *Coordinate `json:"pickupCoordinate"`
}

type IdDist struct {
	Id   int64
	Dist float64
}
