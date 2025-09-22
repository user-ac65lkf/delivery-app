package serv

import (
	"github.com/Shemistan/uzum_delivery/cmd/api/handlers"
	"github.com/Shemistan/uzum_delivery/internal/models"
	"github.com/Shemistan/uzum_delivery/internal/service"
	"github.com/Shemistan/uzum_delivery/internal/storage/postgres"
	desc "github.com/Shemistan/uzum_delivery/pkg/login_v1"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

func GetServ(cnf models.Config, db *sqlx.DB, c desc.LoginV1Client) (*http.Server, error) {

	store := postgres.NewRepoPostgres(db)
	adminService := service.NewService(store, c)
	handler := handlers.NewHandler(adminService)

	router := mux.NewRouter()

	router.HandleFunc("/deliver/v1/healthz", handler.Healthz).Methods(http.MethodGet)
	router.HandleFunc("/deliver/v1/order/{id}", handler.GetOrder).Methods(http.MethodGet)
	router.HandleFunc("/deliver/v1/order", handler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/deliver/v1/order/{id}", handler.CloseOrder).Methods(http.MethodPut)

	srv := &http.Server{
		Addr:        cnf.App.Port,
		ReadTimeout: time.Second * 10,
		Handler:     router,
	}

	return srv, nil
}
