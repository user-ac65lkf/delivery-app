package main

import (
	"database/sql"
	"github.com/Shemistan/uzum_delivery/cmd/conf"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("starting migrate")
	cnf, err := conf.NewConfig()
	if err != nil {
		log.Fatal("failed  to get config:", err.Error())
	}

	sqlConnectionString := conf.GetSqlConnectionString(cnf)
	log.Println(sqlConnectionString)

	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal("failed to open connect with DB", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to get WithInstance", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal("failed to create new Migrate instance", err)
	}
	err = m.Up() // или m.Step(2), если вы хотите явно указать количество запускаемых миграций
	if err != nil {
		log.Fatal("failed to migrate up:", err.Error())
	}

	log.Println("migrate success")
}
