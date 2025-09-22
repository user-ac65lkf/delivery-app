package postgres

import (
	"github.com/Shemistan/uzum_delivery/internal/storage"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewRepoPostgres(db *sqlx.DB) storage.IStorage {
	return &repo{db: db}
}
