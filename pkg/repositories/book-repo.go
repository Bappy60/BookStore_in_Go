package repositories

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DBHandler(db *gorm.DB) {
	DB = db
}

func GetBook(fStruc *types.FilterStruc) ([]models.Book, error) {
	var Books []models.Book

	query := DB.Model(&models.Book{})
	if fStruc.ID == 0 && fStruc.Name == nil && fStruc.Author == nil && fStruc.Publication == nil {
		DB.Find(&Books)
		return Books, nil
	}
	if  fStruc.ID != 0 {
		DB.Where("ID=?", fStruc.ID).Find(&Books)
		return Books, nil
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

func BookCreation(reqStruc *models.Book) (*models.Book, error) {
	DB.Create(reqStruc)
	return reqStruc, nil
}

func UpdateBookInfo(updateBook *models.Book) (*models.Book, error) {

	bookDetails := &models.Book{}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	bookDetails.ID = updateBook.ID

	DB.Save(bookDetails)
	return bookDetails, nil
}

func DeleteBook(ID int64) (string, error) {
	var book = models.Book{}
	DB.Unscoped().Where("ID=?", ID).Delete(book)
	return "Delete successful", nil
}
