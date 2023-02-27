package repositories

import (
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type FilterStruc struct {
	ID          string
	Name        string
	Author      string
	Publication string
}

func GetBook(fStruc *FilterStruc) []models.Book {
	DB = config.GetDB()
	var Books []models.Book
	query := DB.Model(&models.Book{})
	if fStruc.ID == "" && fStruc.Name == "" && fStruc.Author == "" && fStruc.Publication == "" {
		DB.Find(&Books)
		return Books
	}
	if fStruc.ID != "" {
		var getBook models.Book
		ID, err := strconv.ParseInt(fStruc.ID, 0, 0)
		if err != nil {
			return Books
		}
		res := RecordExists(ID)
		if res {
			DB.Where("ID=?", ID).Find(&getBook)
			Books = append(Books, getBook)
			return Books
		}
	}

	if fStruc.Name != "" {
		query = query.Where("name LIKE ?", "%"+fStruc.Name+"%")
	}
	if fStruc.Author != "" {
		query = query.Where("author LIKE ?", "%"+fStruc.Author+"%")
	}
	if fStruc.Publication != "" {
		query = query.Where("publication LIKE ?", "%"+fStruc.Publication+"%")
	}

	if err := query.Find(&Books).Error; err != nil {
		return Books
	}
	return Books
}

func CreateBook(b *models.Book) *models.Book {
	DB = config.GetDB()
	res := DB.NewRecord(b)
	if res {
		exists := RecordExists(b)
		if !exists {
			DB.Create(b)
			return b
		}
	}
	return nil
}

func GetBookByName(bookName string) []models.Book {
	DB = config.GetDB()
	var Books []models.Book
	if bookName == "" {
		return Books
	}
	DB.Where("name=?", bookName).Find(&Books)
	return Books

}

func DeleteBook(Id int64) string {
	DB = config.GetDB()
	var book models.Book
	res := RecordExists(Id)
	if res {
		DB.Unscoped().Where("ID=?", Id).Delete(book)
		return "Delete successful"
	}
	return "There is no book registered by this ID "

}

func RecordExists(arg interface{}) bool {
	var query string
	var args []interface{}
	DB = config.GetDB()
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

	DB.Raw(query, args...).Scan(&result)

	return result.Found
}
