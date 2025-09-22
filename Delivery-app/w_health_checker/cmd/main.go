package main

import (
	"abc/internal/clientdir"
	"abc/internal/initDb"

	_ "github.com/lib/pq"
)

func main() {

	initDb.CreateDbase()
	clientdir.CreateLogs()

}
