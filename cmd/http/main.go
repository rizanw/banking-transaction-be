package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"tx-bank/internal/config"
	"tx-bank/internal/infra"
)

const (
	appName = "tx-bank"
)

func main() {
	var (
		address string
		err     error
	)

	conf, err := config.New(appName)
	if err != nil {
		return
	}

	resources, err := infra.NewResources(conf)
	if err != nil {
		return
	}

	repo := newRepo(resources)
	uc := newUseCase(repo, resources)
	router := newRoutes(uc, resources.SessionManager)

	address = fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	srv := http.Server{
		Addr:         address,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
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
