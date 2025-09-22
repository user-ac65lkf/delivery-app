package handlers

import (
	"encoding/json"
	"github.com/Shemistan/uzum_admin/internal/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) AddProduct(w http.ResponseWriter, r *http.Request) {

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

	res, err := h.serv.AddProduct(ctx, product)
	if err != nil {
		log.Println("failed to add product", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(strconv.FormatInt(res, 10)))
	if err != nil {
		log.Println("failed to write body", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
