package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

var NewBook types.ResponseStruc

var DbHandler = repositories.NewDBHandler(repositories.Initialize())

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Query().Get("bookId")
	bookName := r.URL.Query().Get("bookName")
	bookAuthor := r.URL.Query().Get("author")
	publication := r.URL.Query().Get("publication")

	Fstruc := types.FilterStruc{
		ID:          &bookId,
		Name:        &bookName,
		Author:      &bookAuthor,
		Publication: &publication,
	}

	newBooks, err := DbHandler.GetBook(&Fstruc)
	if err != nil {
		http.Error(w, "Error While getting book from DataBase", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	if len(newBooks) == 0 {
		http.Error(w, "There is no book registered by this ID", http.StatusOK)
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
	newBookStruc := models.Book{
		Name:        CreateBook.Name,
		Author:      CreateBook.Author,
		Publication: CreateBook.Publication,
	}
	err1 := newBookStruc.Validate()
	if err1 != nil {
		http.Error(w, "Invalid Format of Data", http.StatusNotAcceptable)
		return
	}
	book, err := DbHandler.CreateBook(&newBookStruc)
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
	Fstruc := types.FilterStruc{
		ID: &bookId,
	}

	if bookId != "" {
		books, err := DbHandler.GetBook(&Fstruc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
		http.Error(w, "Invalid Format of ID after adding nodemon make file", http.StatusNotAcceptable)
		return
	}
	book := DbHandler.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		w.Write([]byte("Error While Marshaling"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
