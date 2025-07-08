package api

import (
	"log"
	"net/http"
)

func Start() {
	router := NewRouter()

	server := &http.Server{
		Addr:    ":9999",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
