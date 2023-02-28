package repositories

import (
	"errors"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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

func (Dbhandler *DbHandler) GetBook(fStruc *types.FilterStruc) ([]models.Book, error) {
	var Books []models.Book

	query := Dbhandler.db.Model(&models.Book{})
	if fStruc.ID == nil && fStruc.Name == nil && fStruc.Author == nil && fStruc.Publication == nil {
		Dbhandler.db.Find(&Books)
		return Books, nil
	}

	if *fStruc.ID != "" {
		var getBook models.Book

		ID, err := strconv.ParseInt(*fStruc.ID, 0, 0)
		if err != nil {
			return Books, errors.New("invalid format of ID")
		}
		res := Dbhandler.RecordExists(ID)
		if res {
			Dbhandler.db.Where("ID=?", ID).Find(&getBook)
			Books = append(Books, getBook)
			return Books, nil
		}
		return Books,nil
	}

	if fStruc.Name != nil {
		query = query.Where("name LIKE ?", "%"+*fStruc.Name+"%")
	}
	if fStruc.Author != nil {
		query = query.Where("author LIKE ?", "%"+*fStruc.Author+"%")
	}
	if fStruc.Publication != nil {
		query = query.Where("publication LIKE ?", "%"+*fStruc.Publication+"%")
	}

	if err := query.Find(&Books).Error; err != nil {
		return Books, err
	}
	return Books, nil
}

func (Dbhandler *DbHandler) CreateBook(book *models.Book) (*models.Book, error) {
	res := Dbhandler.db.NewRecord(book)
	if res {
		exists := Dbhandler.RecordExists(book)
		if !exists {
			Dbhandler.db.Create(book)
			return book, nil
		}
		return nil, errors.New("book already exists")
	}
	return nil, errors.New("invalid book data")
}

func (Dbhandler *DbHandler) DeleteBook(Id int64) string {
	var book models.Book
	res := Dbhandler.RecordExists(Id)
	if res {
		Dbhandler.db.Unscoped().Where("ID=?", Id).Delete(book)
		return "Delete successful"
	}
	return "There is no book registered by this ID "

}

func (Dbhandler *DbHandler) RecordExists(arg interface{}) bool {
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

	Dbhandler.db.Raw(query, args...).Scan(&result)

	return result.Found
}
