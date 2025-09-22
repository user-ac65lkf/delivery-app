package handlers

import (
	"encoding/json"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := handleToken(h, w, r)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Body.Read", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := &models.Product{}

	err = json.Unmarshal(b, &product)
	if err != nil {
		log.Println("Unmarshal", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	product.ID = int64(id)

	err = h.serv.UpdateProduct(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
