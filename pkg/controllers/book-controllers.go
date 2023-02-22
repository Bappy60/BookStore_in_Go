package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Query().Get("bookId")

	newBooks := models.GetBook(bookId)
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	if len(newBooks) == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("There is no book registered by this ID"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CBook := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&CBook)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	newBookStruc := models.Book{
		Name:        CBook.Name,
		Author:      CBook.Author,
		Publication: CBook.Publication,
	}
	err1 := newBookStruc.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}
	b := newBookStruc.CreateBook()
	if b != nil {
		res, err := json.Marshal(b)
		if err != nil {
			http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte("Book already created"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBook = &models.Book{}
	err := json.NewDecoder(r.Body).Decode(updateBook)
	if err != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}
	bookId := r.URL.Query().Get("bookId")

	if bookId != "" {
		books := models.GetBook(bookId)
		if books != nil {

			bookDetails := books[0]
			db := config.GetDB()

			if updateBook.Name != "" {
				bookDetails.Name = updateBook.Name
			}
			if updateBook.Author != "" {
				bookDetails.Author = updateBook.Author
			}
			if updateBook.Publication != "" {
				bookDetails.Publication = updateBook.Publication
			}

			db.Save(&bookDetails)
			res, err := json.Marshal(bookDetails)
			if err != nil {
				http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
				return
			}
			w.WriteHeader(http.StatusAccepted)
			w.Write(res)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("ID not Found"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Format of ID", http.StatusNotAcceptable)
		return
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		w.Write([]byte("Error While Marshaling"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
