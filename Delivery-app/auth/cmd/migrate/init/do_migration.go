package migration

import (
	"database/sql"
	"github.com/Shemistan/uzum_auth/cmd/conf"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func DoMigration() {
	cnf, err := conf.NewConfig()
	if err != nil {
		log.Fatal("failed  to get config:", err.Error())
	}

	sqlConnectionString := conf.GetSqlConnectionString(cnf)

	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal("failed to open connect with DB", err.Error())
	}

	var isThere bool

	row := db.QueryRow(`SELECT EXISTS (
								SELECT FROM
								pg_tables
								WHERE
								tablename = 'users');`)
	err = row.Scan(&isThere)
	if err != nil {
		log.Println("Error from auth/migration ", err)
	}

	if !isThere {
		log.Println("starting migrate")

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

		log.Println("auth/users table created")
	}
}
