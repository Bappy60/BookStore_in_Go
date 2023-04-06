package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

// var BookService domain.IBookService

// func SetBookService(bService domain.IBookService) {
// 	BookService = bService
// }
type BookController struct {
	bookService domain.IBookService
}

func BookControllerInstance(bookService domain.IBookService) domain.IBookController {
	return &BookController{
		bookService: bookService,
	}
}

func (bookController *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Query().Get("bookId")
	bookName := r.URL.Query().Get("bookName")
	numOfPages := r.URL.Query().Get("number_of_pages")
	authorID := r.URL.Query().Get("author_id")
	publication := r.URL.Query().Get("publication")
	publicationYear := r.URL.Query().Get("publication_year")

	reqStruc := types.BookReqStruc{
		ID:              bookId,
		Name:            bookName,
		NumberOfPages:   numOfPages,
		AuthorID:        authorID,
		Publication:     publication,
		PublicationYear: publicationYear,
	}

	newBooks, err := bookController.bookService.GetBooks(&reqStruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	if len(newBooks) == 0 {
		http.Error(w, "no book registered ", http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (bookController *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := types.CreateBookStruc{}
	err := json.NewDecoder(r.Body).Decode(&CreateBook)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}

	err1 := CreateBook.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}

	book, err := bookController.bookService.CreateBook(&CreateBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func (bookController *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBookStruc = &types.UpdateBookStruc{}
	err := json.NewDecoder(r.Body).Decode(updateBookStruc)
	if err != nil {
		http.Error(w, "Invalid Format of Data while decoding", http.StatusNotAcceptable)
		return
	}
	err1 := updateBookStruc.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data while validating", http.StatusNotAcceptable)
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	updateBookStruc.ID = bookId

	updatedBook, err := bookController.bookService.UpdateBook(updateBookStruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func (bookController *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if bookId == "" {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	msg, err := bookController.bookService.DeleteBook(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
}
