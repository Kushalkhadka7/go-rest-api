package main

import (
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/routes"
	"go-rest-api/server"
	"log"
	"os"
)

// main app entry point.
func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := db.New(config)
	if err != nil {
		panic(err)
	}
	db.Migrate()

	router := routes.New(db)

	srv := server.New(router, config)
	if err := srv.Start(); err != nil {
		log.Fatalf("Error on starting the server %s", err)
		os.Exit(1)
	}

	defer db.Close()
}
