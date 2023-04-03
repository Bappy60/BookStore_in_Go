package routes

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/author/{author_id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{author_id}", controllers.DeleteAuthor).Methods("DELETE")
}
