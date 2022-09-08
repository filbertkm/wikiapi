// Package Wikipedia descriptions API
//
// Wikipedia descriptions API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/filbertkm/wikiapi/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "[API]", log.LstdFlags)
	dh := handlers.NewDescription(l)

	sm := mux.NewRouter()

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	sm.Handle("/docs", sh).Methods("GET")
	sm.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	sm.HandleFunc("/page/{title}", dh.GetDescription).Methods("GET")

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}