package models

import (
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&b.Author, validation.Required, validation.Length(3, 50)),
		validation.Field(&b.Publication),
	)
}

func Initialize() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	res := db.NewRecord(b)
	if res {
		exists := AlreadyExists(b)
		if !exists {
			db.Create(&b)
			return b
		}
	}
	return nil
}

func GetBook(bookId string) []Book {
	var Books []Book
	if bookId == "" {
		db.Find(&Books)
		return Books
	}
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return Books
	}

	var getBook Book
	res := IsValid(ID)
	if res {
		db.Where("ID=?", ID).Find(&getBook)
		Books = append(Books, getBook)
		return Books
	}
	return Books

}

func DeleteBook(Id int64) string {

	var book Book
	res := IsValid(Id)
	if res {
		db.Unscoped().Where("ID=?", Id).Delete(book)
		return "Delete successful"
	}
	return "There is no book registered by this ID "

}

func IsValid(Id int64) bool {

	var result struct {
		Found bool
	}

	db.Raw("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?) AS found", Id).Scan(&result)

	return result.Found
}
func AlreadyExists(b *Book) bool {
	var result struct {
		Found bool
	}

	db.Raw("SELECT EXISTS(SELECT 1 FROM books WHERE name = ? AND author = ?) AS found", b.Name, b.Author).Scan(&result)

	return result.Found
}
