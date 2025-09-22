package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	ctx, _ := handleToken(h, w, r)

	id := mux.Vars(r)["id"]

	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.serv.CloseOrder(ctx, int64(intID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
