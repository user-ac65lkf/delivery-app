package storage

import (
	"abc/internal/models"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

type IStorage interface {
	InsertData(fromChannel []*models.CheckStruct)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) IStorage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) InsertData(fromChannel []*models.CheckStruct) {

	txn, err := s.db.Begin()
	if err != nil {
		log.Fatalf("Could not start psql transaction: %s", err.Error())
	}

	stmt, err := txn.Prepare(pq.CopyIn("checktable", "datetime", "server", "status"))
	if err != nil {
		err := txn.Rollback()
		if err != nil {
			return
		}
		log.Fatalf("Could not prepare psql statement: %s", err.Error())
	}

	// Iterate the data and add the data to the psql statement
	for _, v := range fromChannel {
		_, err := stmt.Exec(v.Datetime, v.Server, v.Status)
		if err != nil {
			log.Fatalf("Could not execute the statement: %s", err)
		}
	}

	// Execute, commit, and close the transaction
	err = stmt.Close()
	if err != nil {
		log.Fatalf("Could not close the psql statement: %s", err.Error())
	}

	err = txn.Commit()
	if err != nil {
		log.Fatalf("Could not commit the psql transaction: %s", err.Error())
	}

	fmt.Printf("\nlog has been successfully inserted.\n Open 'checktable' to see all logs. \nNext batch of logs will be inserted after 12 seconds")
}
