package repositories

import (
	"errors"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DBHandler(db *gorm.DB) {
	DB = db
}

func GetBooks(filterStruct *types.FilterStruc) ([]models.Book, error) {
	var Books []models.Book

	query := DB.Model(&models.Book{}).Preload("Author")
	if filterStruct.ID == 0 && *filterStruct.Name == "" && *filterStruct.AuthorID == 0 && *filterStruct.Publication == "" && filterStruct.NumberOfPages == 0 && filterStruct.PublicationYear == 0 {
		query.Find(&Books)
		return Books, nil
	}
	if filterStruct.ID != 0 {
		query.Where("ID=?", filterStruct.ID).Find(&Books)
		return Books, nil
	}
	if *filterStruct.Name != "" {
		query = query.Where("name LIKE ?", "%"+*filterStruct.Name+"%")
	}
	if *filterStruct.AuthorID != 0 {
		query = query.Where("author_id = ?", *filterStruct.AuthorID)
	}
	if *filterStruct.Publication != "" {
		query = query.Where("publication LIKE ?", *filterStruct.Publication)
	}
	if filterStruct.PublicationYear != 0 {
		query = query.Where("publication_year = ?", filterStruct.PublicationYear)
	}
	if filterStruct.NumberOfPages != 0 {
		query = query.Where("number_of_pages = ?", filterStruct.NumberOfPages)
	}
	if err := query.Find(&Books).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return Books, err
	}

	return Books, nil
}

func BookCreation(bookStruc *models.Book) (*models.Book, error) {
	var existingBook models.Book
	err := DB.Where("name = ? AND author_id = ?", bookStruc.Name, bookStruc.AuthorID).First(&existingBook).Error
	if err == nil {
		return &models.Book{}, errors.New("book already exists")
	}
	if err := DB.Create(bookStruc).Error; err != nil {
		return nil, DB.Error
	}
	return bookStruc, nil
}

func UpdateBookInfo(updateBook *types.UpdateBookStruc) (*models.Book, error) {

	bookDetails := &models.Book{}
	bookDetails.ID = updateBook.ID
	if err := DB.Where("id = ?", bookDetails.ID).Find(bookDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no book found with given ID")
		}
	}
	if updateBook.Name != nil {
		bookDetails.Name = *updateBook.Name
	}
	if updateBook.NumberOfPages != nil {
		bookDetails.NumberOfPages = *updateBook.NumberOfPages
	}
	if updateBook.PublicationYear != nil {
		bookDetails.PublicationYear = *updateBook.PublicationYear
	}
	if updateBook.Publication != nil {
		bookDetails.Publication = *updateBook.Publication
	}

	DB.Save(bookDetails)
	return bookDetails, nil
}

func DeleteBook(ID int64) (string, error) {
	var book = models.Book{}
	DB.Unscoped().Where("ID=?", ID).Delete(book)
	return "Delete successful", nil
}
