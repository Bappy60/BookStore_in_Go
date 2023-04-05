package repositories

import (
	"errors"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

func (repo *bookRepo) GetBooks(filterStruct *types.FilterBookStruc) ([]models.Book, error) {
	var Books []models.Book

	query := repo.db.Model(&models.Book{}).Preload("Author")
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

func (repo *bookRepo) CreateBook(newBook *types.CreateBookStruc) (*models.Book, error) {
	var Book models.Book
	err := repo.db.Where("name = ? AND author_id = ?", newBook.Name, newBook.AuthorID).First(&Book).Error
	if err == nil {
		return nil, errors.New("book already exists")
	}

	Book.Name = newBook.Name
	Book.Publication = newBook.Publication
	Book.PublicationYear = newBook.PublicationYear
	Book.AuthorID = newBook.AuthorID
	Book.NumberOfPages = newBook.NumberOfPages

	if err := repo.db.Create(&Book).Error; err != nil {
		return nil, repo.db.Error
	}
	return &Book, nil
}

func (repo *bookRepo) UpdateBook(updateBook *models.Book) (*models.Book, error) {
	 bookDetails := &models.Book{}
	if err := repo.db.Where("id = ?", updateBook.ID).Find(bookDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no book found with given ID")
		}
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.NumberOfPages != 0 {
		bookDetails.NumberOfPages = updateBook.NumberOfPages
	}
	if updateBook.PublicationYear != 0 {
		bookDetails.PublicationYear = updateBook.PublicationYear
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	repo.db.Save(bookDetails)
	return bookDetails, nil
}

func (repo *bookRepo) DeleteBook(bookID int64) (string, error) {
	var book = models.Book{}
	if err := repo.db.Where("id = ?", bookID).Find(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("no book found with given ID")
		}
	}
	repo.db.Unscoped().Where("ID=?", bookID).Delete(book)
	return "Delete successful", nil
}