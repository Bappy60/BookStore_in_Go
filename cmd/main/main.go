package main

import (
	"log"
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	models.Initialize()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server Started...")
	log.Fatal(http.ListenAndServe("localhost:9011", r))
	
}
