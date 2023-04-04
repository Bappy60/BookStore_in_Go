package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

var BookRepo domain.IBookRepo

func SetBookRepo(bRepo domain.IBookRepo) {
	BookRepo = bRepo
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Query().Get("bookId")
	bookName := r.URL.Query().Get("bookName")
	NumOfPages := r.URL.Query().Get("number_of_pages")
	AuthorID := r.URL.Query().Get("author_id")
	publication := r.URL.Query().Get("publication")
	PublicationYear := r.URL.Query().Get("publication_year")

	parsedId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil && bookId != "" {
		http.Error(w, "invalid format of book id", http.StatusBadRequest)
		return
	}
	parsedAuthorID, err := strconv.ParseUint(AuthorID, 0, 0)
	if err != nil && AuthorID != "" {
		http.Error(w, "invalid format of author id", http.StatusBadRequest)
		return
	}
	parsedNumOfPages, err := strconv.ParseInt(NumOfPages, 0, 0)
	if err != nil && NumOfPages != "" {
		http.Error(w, "invalid format of number of pages", http.StatusBadRequest)
		return
	}
	parsedPublicationYear, err := strconv.ParseInt(PublicationYear, 0, 0)
	if err != nil && PublicationYear != "" {
		http.Error(w, "invalid format of publication year", http.StatusBadRequest)
		return
	}

	pAuthorID := uint(parsedAuthorID)
	Fstruc := types.FilterStruc{
		ID:              uint(parsedId),
		Name:            &bookName,
		PublicationYear: int(parsedPublicationYear),
		NumberOfPages:   int(parsedNumOfPages),
		AuthorID:        &pAuthorID,
		Publication:     &publication,
	}

	newBooks, err := BookRepo.GetBooks(&Fstruc)
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
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

	book, err := BookRepo.CreateBook(&CreateBook)
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBook = &types.UpdateBookStruc{}
	err := json.NewDecoder(r.Body).Decode(updateBook)
	if err != nil {
		http.Error(w, "Invalid Format of Data while decoding", http.StatusNotAcceptable)
		return
	}
	err1 := updateBook.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data while validating", http.StatusNotAcceptable)
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	bookID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	updateBook.ID = uint(bookID)
	updatedBook, err := BookRepo.UpdateBook(updateBook)
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	if bookId == "" {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	parsedId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	msg, err := BookRepo.DeleteBook(parsedId)
	if err != nil {
		http.Error(w, "There is no book registered by this ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
}
