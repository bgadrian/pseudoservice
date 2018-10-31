package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"

	"github.com/bgadrian/pseudoservice/server"
)

type config struct {
	//TODO make user system with APIKEYS and Quotas
	ApiKey   string `env:"APIKEY" envDefault:"SECRET42"`
	BasePath string `env:"BASEPATH" envDefault:"/api/v1"`
	Port     int    `env:"PORT" envDefault:"8080"`
	Debug    bool   `env:"DEBUG" envDefault:"false"`
}

func main() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("error ENV variables: %v", err)
	}

	//for --help or -h
	if len(os.Args) > 1 {
		fmt.Println("Accepting the following parameters:")
		fmt.Println("ENV variable: PORT")
		fmt.Println("ENV variable: APIKEY")
		fmt.Println("ENV variable: DEBUG")
		fmt.Println("HTTP query: ?token=MYAPIKEY")
		fmt.Println("HTTP header: Accept-Encoding: gzip")
		os.Exit(0)
	}

	routes := server.DefaultRoutes(cfg.BasePath)
	middlewares := []mux.MiddlewareFunc{}
	apiKeeper := server.NewApiKeyHandler([]string{cfg.ApiKey})
	middlewares = append(middlewares, apiKeeper.Handle)
	middlewares = append(middlewares, server.Gzip, server.CustomHeaders)
	if cfg.Debug {
		middlewares = append(middlewares, server.Logger)
	}

	router := server.NewRouter(routes, middlewares)
	srv := server.New(router, ":"+strconv.Itoa(cfg.Port))

	go func() { log.Fatal(srv.ListenAndServe()) }()
	log.Printf("server started on 0.0.0.0:%d\n", cfg.Port)
	log.Printf("waiting for an OS Interrupt signal (Ctrl-C) ...\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	log.Println("shutting down signal received, waiting ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}
