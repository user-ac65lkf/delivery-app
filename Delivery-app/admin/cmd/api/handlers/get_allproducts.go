package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := handleToken(h, w, r)
	// get all the stocks in the db
	products, err := h.serv.GetAllProducts(ctx)

	if err != nil {
		log.Fatalf("Unable to get all stock. %v", err)
	}

	data, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("Unable to get all stock. %v", err)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
