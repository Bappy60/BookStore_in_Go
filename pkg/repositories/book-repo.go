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

func Initialize() *gorm.DB {
	config.Connect()
	db := config.GetDB()
	return db
}

type DbHandler struct {
	db *gorm.DB
}

func NewDBHandler(db *gorm.DB) *DbHandler {
	return &DbHandler{
		db: db,
	}
}

func (h *DbHandler) GetBook(fStruc *FilterStruc) []models.Book {
	var Books []models.Book
	query := h.db.Model(&models.Book{})
	if fStruc.ID == "" && fStruc.Name == "" && fStruc.Author == "" && fStruc.Publication == "" {
		h.db.Find(&Books)
		return Books
	}
	if fStruc.ID != "" {
		var getBook models.Book
		ID, err := strconv.ParseInt(fStruc.ID, 0, 0)
		if err != nil {
			return Books
		}
		res := h.RecordExists(ID)
		if res {
			h.db.Where("ID=?", ID).Find(&getBook)
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

func (h *DbHandler) CreateBook(b *models.Book) *models.Book {
	res := h.db.NewRecord(b)
	if res {
		exists := h.RecordExists(b)
		if !exists {
			h.db.Create(b)
			return b
		}
	}
	return nil
}

func (h *DbHandler) DeleteBook(Id int64) string {
	var book models.Book
	res := h.RecordExists(Id)
	if res {
		h.db.Unscoped().Where("ID=?", Id).Delete(book)
		return "Delete successful"
	}
	return "There is no book registered by this ID "

}

func (h *DbHandler) RecordExists(arg interface{}) bool {
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

	h.db.Raw(query, args...).Scan(&result)

	return result.Found
}
