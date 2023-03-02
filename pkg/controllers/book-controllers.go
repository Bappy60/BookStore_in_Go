package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

var NewBook types.ResponseStruc

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Query().Get("bookId")
	bookName := r.URL.Query().Get("bookName")
	bookAuthor := r.URL.Query().Get("author")
	publication := r.URL.Query().Get("publication")
	parsedId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil && bookId != "" {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	Fstruc := types.FilterStruc{
		ID:          uint(parsedId),
		Name:        &bookName,
		Author:      &bookAuthor,
		Publication: &publication,
	}

	newBooks, err := repositories.GetBook(&Fstruc)
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
	CreateBook := models.Book{}
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
	Fstruc := types.FilterStruc{
		Name:        &CreateBook.Name,
		Author:      &CreateBook.Author,
		Publication: &CreateBook.Publication,
	}

	newBooks, err := repositories.GetBook(&Fstruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(newBooks) == 0 {
		reqStruct := models.Book{
			Name:        CreateBook.Name,
			Author:      CreateBook.Author,
			Publication: CreateBook.Publication,
		}
		book, err := repositories.BookCreation(&reqStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if book != nil {
			res, err := json.Marshal(book)
			if err != nil {
				http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("book already exists"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBook = &models.Book{}
	err := json.NewDecoder(r.Body).Decode(updateBook)
	if err != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}
	err1 := updateBook.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	parsedId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	Fstruc := types.FilterStruc{
		ID: uint(parsedId),
	}
	books, err := repositories.GetBook(&Fstruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(books) == 0 {
		http.Error(w, "there is no book registered by this ID", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	updateBook.ID = uint(bookID)
	updatedBook, err := repositories.UpdateBookInfo(updateBook)
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
	Fstruc := types.FilterStruc{
		ID: uint(parsedId),
	}
	books, err := repositories.GetBook(&Fstruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(books) == 0 {
		http.Error(w, "there is no book registered by this ID", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	msg, err := repositories.DeleteBook(parsedId)
	if err != nil {
		http.Error(w, "There is no book registered by this ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
}
