package handlers

import (
	"net/http"

	"github.com/Shemistan/uzum_delivery/internal/service"
)

type IHandlers interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	CloseOrder(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{serv: serv}
}

type handler struct {
	serv service.IService
}
