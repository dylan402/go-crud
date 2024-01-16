package main

import (
	"fmt"
	"go-crud/configs"
	"go-crud/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", handlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/{id:[0-9]+}", handlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/", handlers.List).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", handlers.Get).Methods(http.MethodGet)

	http.Handle("/", router)

	port := fmt.Sprintf(":%s", configs.GetApiConfig().Port)
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
