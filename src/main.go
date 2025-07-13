package main

import (
	"log"

	"github.com/mathesukkj/rinha-de-backend-2025/src/api"
	"github.com/mathesukkj/rinha-de-backend-2025/src/db"
)

func main() {
	db, err := db.LoadDB()
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	api := api.NewAPI(db)
	api.Start()
}
