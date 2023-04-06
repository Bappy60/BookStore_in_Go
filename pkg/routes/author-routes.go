package routes

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/gorilla/mux"
)

func AuthorRoutes(router *mux.Router, authorController domain.IAuthorController) {
	router.HandleFunc("/author", authorController.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", authorController.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{author_id}", authorController.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{author_id}", authorController.DeleteAuthor).Methods("DELETE")
}
