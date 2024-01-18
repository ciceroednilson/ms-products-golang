package main

import (
	"go/core/usercases/productusercase"
	"go/handlers/producthandler"
	"go/infrastructure/repositories/productrepository"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	productrepository := productrepository.NewRepository()
	productusercase := productusercase.NewProductUserCase(productrepository)
	producthandler := producthandler.NewProductHandler(productusercase)
	startServer(producthandler)
}

func startServer(producthandler producthandler.ProductHandler) {
	router := mux.NewRouter()
	sub := router.PathPrefix("/products").Subrouter()
	sub.HandleFunc("/{key}", producthandler.Read).Methods("GET")
	sub.HandleFunc("/{key}/details", producthandler.ReadDetails).Methods("GET")
	sub.HandleFunc("", producthandler.ReadAll).Methods("GET")
	sub.HandleFunc("", producthandler.Create).Methods("POST")
	sub.HandleFunc("", producthandler.Update).Methods("PUT")
	sub.HandleFunc("/{key}", producthandler.Delete).Methods("DELETE")

	server := &http.Server{
		Handler:      sub,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         "127.0.0.1:8000",
	}
	log.Fatal(server.ListenAndServe())
}
