package container

import (
	"log"
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/connection"
	"github.com/Bappy60/BookStore_in_Go/pkg/controllers"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/routes"
	"github.com/Bappy60/BookStore_in_Go/pkg/services"
	"github.com/gorilla/mux"
)

func Serve() {
	config.SetConfig()
	var db = connection.Initialize()

	bookRepo := repositories.BookDBInstance(db)
	bookService := services.BookServiceInstance(bookRepo)
	bookController := controllers.BookControllerInstance(bookService)

	authorRepo := repositories.AuthorDBInstance(db)
	authorService := services.AuthorServiceInstance(authorRepo)
	authorController := controllers.AuthorControllerInstance(authorService)

	log.Println("Database Connected...")
	r := mux.NewRouter()
	routes.AuthorRoutes(r, authorController)
	routes.BookRoutes(r,bookController)
	http.Handle("/", r)
	log.Println("Server Started...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
