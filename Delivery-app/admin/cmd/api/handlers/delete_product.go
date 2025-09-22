package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := handleToken(h, w, r)

	params := mux.Vars(r)

	productId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	err = h.serv.DeleteProduct(ctx, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
