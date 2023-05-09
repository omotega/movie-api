package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"grans/pkg/config"
)

type application struct {
	logger *log.Logger
}


func main() {
	appConfig,err := configs.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	logger := log.New(os.Stdout,"",log.Ldate | log.Ldate)

	app := &application {
		logger: logger,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d",appConfig.Port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s",appConfig.Env,srv.Addr)
	error := srv.ListenAndServe()
	logger.Fatal(error)
}