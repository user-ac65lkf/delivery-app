package handlers

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/service"
	"net/http"
)

type IHandlers interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{serv: serv}
}

type handler struct {
	serv service.IService
}

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

	w.WriteHeader(http.StatusOK)
}
