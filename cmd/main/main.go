package main

import (
	"log"
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	var db = config.Initialize()
	db.AutoMigrate(&models.Book{})
	repositories.DBHandler(db)
	log.Println("Database Connected...")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server Started...")
	log.Fatal(http.ListenAndServe("localhost:9012", r))
}
