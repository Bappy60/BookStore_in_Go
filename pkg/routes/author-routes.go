package routes

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/controllers"
	"github.com/gorilla/mux"
)
func AuthorRoutes(router *mux.Router) {
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{author_id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{author_id}", controllers.DeleteAuthor).Methods("DELETE")
}