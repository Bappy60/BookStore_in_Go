package main

import (
	"log"
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/controllers"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/routes"
	"github.com/Bappy60/BookStore_in_Go/pkg/services"
	"github.com/gorilla/mux"
)

func main() {
	var db = config.Initialize()

	bookRepo := repositories.BookDBInstance(db)
	bookService := services.BookServiceInstance(bookRepo)
	controllers.SetBookService(bookService)

	authorRepo := repositories.AuthorDBInstance(db)
	authorService := services.AuthorServiceInstance(authorRepo)
	controllers.SetAuthorService(authorService)

	db.AutoMigrate(&models.Book{},&models.Author{})
	log.Println("Database Connected...")
	r := mux.NewRouter()
	routes.AuthorRoutes(r)
    routes.BookRoutes(r)
	http.Handle("/", r)
	log.Println("Server Started...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
