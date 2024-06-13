package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		address = "0.0.0.0:8080"
		err     error
	)

	_ = newRepo()
	router := newRoutes()

	srv := http.Server{
		Addr:         address,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	log.Println("app starting on:", address)
	err = srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
