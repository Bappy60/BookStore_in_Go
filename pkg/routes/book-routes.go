package routes

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/gorilla/mux"
)

var BookRoutes = func(router *mux.Router, bookController domain.IBookController) {

	router.HandleFunc("/book", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books", bookController.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bookController.DeleteBook).Methods("DELETE")

}
