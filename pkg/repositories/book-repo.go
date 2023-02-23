package repositories

import (
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CreateBook(b *models.Book) *models.Book {
	db = config.GetDB()
	res := db.NewRecord(b)
	if res {
		exists := RecordExists(b)
		if !exists {
			db.Create(b)
			return b
		}
	}
	return nil
}

func GetBook(bookId string) []models.Book {
	db = config.GetDB()
	var Books []models.Book
	if bookId == "" {
		db.Find(&Books)
		return Books
	}
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return Books
	}

	var getBook models.Book
	res := RecordExists(ID)
	if res {
		db.Where("ID=?", ID).Find(&getBook)
		Books = append(Books, getBook)
		return Books
	}
	return Books

}

func DeleteBook(Id int64) string {

	db = config.GetDB()
	var book models.Book
	res := RecordExists(Id)
	if res {
		db.Unscoped().Where("ID=?", Id).Delete(book)
		return "Delete successful"
	}
	return "There is no book registered by this ID "

}

func RecordExists(arg interface{}) bool {
	var query string
	var args []interface{}

	switch v := arg.(type) {
	case int64:
		query = "SELECT EXISTS(SELECT 1 FROM books WHERE id = ?) AS found"
		args = []interface{}{v}
	case *models.Book:
		query = "SELECT EXISTS(SELECT 1 FROM books WHERE name = ? AND author = ?) AS found"
		args = []interface{}{v.Name, v.Author}
	default:
		return false
	}

	var result struct {
		Found bool
	}

	db.Raw(query, args...).Scan(&result)

	return result.Found
}
