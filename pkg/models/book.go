package models

import (
	"log"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Initialize() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	log.Printf("One New Book named : %s is Created", b.Name)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	var result struct {
		Found bool
	}
	db.Raw("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?) AS found", Id).Scan(&result)

	if result.Found {
		db := db.Where("ID=?", Id).Find(&getBook)
		return &getBook, db
	} else {
		// does not exist
		return nil, nil
	}
}

func DeleteBook(ID int64) string {
	var book Book
	var result struct {
		Found bool
	}

	db.Raw("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?) AS found", ID).Scan(&result)

	if result.Found {
		// exists
		db.Unscoped().Where("ID=?", ID).Delete(book)
		return "Delete successful"
	} else {
		// does not exist
		return "There is no book registered by this ID "
	}

}
