package initDb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
)

var (
	createDBQuery        = "CREATE DATABASE %s;"
	createUserQuery      = `CREATE USER %s with PASSWORD 'checkh';`
	grantPrivilegesQuery = "GRANT ALL ON DATABASE %s TO %s;"
)

func CreateDbase() {
	dbNewInfo := "postgres://postgres:postgres@localhost:5432/checkh?sslmode=disable"

	dbNew, err := sql.Open("postgres", dbNewInfo)
	if err != nil {
		log.Println("dbNew", err)
	}

	err1 := dbNew.Ping()
	if err1 != nil {
		dsn := "postgres://postgres:postgres@localhost:5432/service?sslmode=disable"
		dbName := "checkh"
		dbUser := "checkh"

		log.Println("Creating db and user")

		db, err := sql.Open("postgres", dsn)
		err = db.Ping()

		if err != nil {
			log.Fatal(errors.Wrapf(err, "Could not open sql connection to %s", dsn))
		}
		defer db.Close()
		fmt.Println(".....>> ", fmt.Sprintf(createUserQuery, dbUser))

		//// Создаем пользователя
		var hasRole string

		row := db.QueryRow(`SELECT 1 FROM pg_roles WHERE rolname='checkh';`)

		err = row.Scan(&hasRole)
		if err != nil {
			fmt.Println(err)
			if _, err := db.Exec(fmt.Sprintf(createUserQuery, dbUser)); err != nil {
				log.Fatal(errors.Wrapf(err, "Could not execute query %s with params %s", createUserQuery, dbUser))
			}
		}

		// Создаем базу данных
		if _, err := db.Exec(fmt.Sprintf(createDBQuery, dbName)); err != nil {
			log.Fatal(errors.Wrapf(err, "Could not create database %s", dbName))
		}

		// Грантим привилегии
		if _, err := db.Exec(fmt.Sprintf(grantPrivilegesQuery, dbName, dbUser)); err != nil {
			log.Fatal(errors.Wrapf(err, "Could not execute query %s with params %s %s", grantPrivilegesQuery, dbName, dbUser))
		}

		if _, err := db.Exec(fmt.Sprintf(`ALTER DATABASE checkh owner TO checkh;`)); err != nil {
			log.Fatal(errors.Wrapf(err, "Could not ALTER DATABASE checkh owner TO checkh;"))
		}

		log.Println("DB and user creation finished")

		dbNewString := "postgres://checkh:checkh@localhost:5432/checkh?sslmode=disable"

		dbNewConn, err := sql.Open("postgres", dbNewString)

		if err != nil {
			log.Println("dbNewString", err)
		}

		res, err := dbNewConn.Exec(`CREATE TABLE IF NOT EXISTS checktable (
								id serial primary key,
								datetime varchar(255),
    							server varchar(255),
								status varchar(20)    							
	               				);`)
		if err != nil {
			log.Fatal("Failed to create table", err)
		}

		log.Printf("table created %+v", res)
	}
}
