package main

import (
	"context"
	"flag"
	migration "github.com/Shemistan/uzum_auth/cmd/migrate/init"
	"github.com/Shemistan/uzum_auth/internal/app"
	"log"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("failed to create app")
	}

	migration.DoMigration()

	err = a.Run()
	if err != nil {
		log.Fatal("failed to run app")
	}
}
