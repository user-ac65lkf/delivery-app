package handlers

import (
	"net/http"
)

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
