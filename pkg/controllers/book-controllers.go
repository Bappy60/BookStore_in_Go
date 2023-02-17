package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		r, _ := json.Marshal("Invalid Format of ID")
		w.Header().Set("Content-Type", "application/json")
		w.Write(r)
	} else {
		bookDetails, _ := models.GetBookById(ID)
		if bookDetails != nil {
			res, _ := json.Marshal(bookDetails)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			res, _ := json.Marshal("There is no book registered by this ID")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(res)
		}

	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	if CreateBook.Name == "" || CreateBook.Author == "" || CreateBook.Publication == "" {
		r, _ := json.Marshal("Name, Author and Publication can't be null ")
		
		w.Write(r)
	} else {
		b := CreateBook.CreateBook()
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		r, _ := json.Marshal("Invalid Format of ID")
		w.Header().Set("Content-Type", "application/json")
		w.Write(r)
	} else {
		book := models.DeleteBook(ID)
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		r, _ := json.Marshal("Invalid Format of ID")
		w.Header().Set("Content-Type", "application/json")
		w.Write(r)
	} else {
		bookDetails, db := models.GetBookById(ID)
		if bookDetails != nil {

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
			res, _ := json.Marshal(bookDetails)
			w.Header().Set("Content-Type", " application/json")
			w.WriteHeader(200)
			w.Write(res)
		} else {
			res, _ := json.Marshal("There is no book registered by this ID")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(res)
		}
	}
}
