package handlers

import (
	"encoding/json"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx, _ := handleToken(h, w, r)
	longitude := r.URL.Query().Get("longitude")
	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	latitude := r.URL.Query().Get("latitude")
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	courierCoordinate := &models.Coordinate{
		Longitude: long,
		Latitude:  lat,
	}

	res, err := h.serv.GetOrders(ctx, courierCoordinate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("failed to write in body")
	}
}
