package api

import (
	"log"
	"net/http"

	"github.com/mathesukkj/rinha-de-backend-2025/src/db"
)

type API struct {
	db *db.DB
}

func NewAPI(db *db.DB) *API {
	return &API{db: db}
}

func (api *API) Start() {
	router := NewRouter(api)

	server := &http.Server{
		Addr:    ":9999",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
