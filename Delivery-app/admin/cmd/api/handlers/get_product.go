package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := handleToken(h, w, r)

	params := mux.Vars(r)

	productId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	product, err := h.serv.GetProduct(ctx, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Unable to get product. %v", err)
	}

	w.Header().Add("content-type", "application/json")

	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Unable to write data. %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
