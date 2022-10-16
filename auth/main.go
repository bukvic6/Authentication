package main

import (
	"Microservices/auth/handlers"
	"Microservices/auth/middleware"
	"Microservices/homehandlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := mux.NewRouter()
	s := r.Methods(http.MethodPost).Subrouter()
	s.HandleFunc("/login", handlers.Login)
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := homehandlers.NewProducts(l)

	m := r.PathPrefix("auth").Subrouter()
	m.HandleFunc("/", ph.GetProducts).Methods(http.MethodGet)
	m.Use(middleware.VerifyJwt)
	srv := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Server starting on port 9090")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Service shutting down...")
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")

}
