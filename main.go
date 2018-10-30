package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/bgadrian/pseudoservice/server"
)

func main() {
	port := "8080"
	key := "SECRET42"
	basePath := "/api/v1" //NO end slash

	routes := server.DefaultRoutes(basePath)
	middlewares := []mux.MiddlewareFunc{server.Logger}
	apiKeeper := server.NewApiKeyHandler([]string{key})
	middlewares = append(middlewares, apiKeeper.Handle)

	router := server.NewRouter(routes, middlewares)
	srv := server.New(router, ":"+port)

	go func() { log.Fatal(srv.ListenAndServe()) }()
	log.Printf("server started on :%s", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	log.Println("shutting down signal received, waiting ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}
