package serv

import (
	"github.com/Shemistan/uzum_admin/cmd/api/handlers"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/service"
	"github.com/Shemistan/uzum_admin/internal/storage/postgres"
	desc "github.com/Shemistan/uzum_admin/pkg/login_v1"
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

	router.HandleFunc("/admin/v1/healthz", handler.Healthz).Methods(http.MethodGet)
	router.HandleFunc("/product/add", handler.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", handler.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", handler.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", handler.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/product/{id}", handler.DeleteProduct).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:        cnf.App.Port,
		ReadTimeout: time.Second * 10,
		Handler:     router,
	}

	return srv, nil
}
