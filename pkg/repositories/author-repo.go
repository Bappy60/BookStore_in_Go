package repositories

import (
	"errors"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

func GetAuthor(authorStruc *types.AuthorStruc) ([]models.Author, error) {
	var Authors []models.Author

	query := DB.Model(&models.Author{}).Preload("Books")
	if authorStruc.ID == 0 && authorStruc.Name == "" && authorStruc.Email == "" && authorStruc.Age == 0 {
		query.Find(&Authors)
		return Authors, nil
	}

	if authorStruc.ID != 0 {
		query.Where("id =?", authorStruc.ID).Find(&Authors)
		return Authors, nil
	} else {
		if authorStruc.Name != "" {
			query = query.Where("name LIKE ?", "%"+authorStruc.Name+"%")
		}
		if authorStruc.Email != "" {
			query = query.Where("email LIKE ?", "%"+authorStruc.Email+"%")
		}
		if authorStruc.Age != 0 {
			query = query.Where("age =?", authorStruc.Age)
		}
		if err := query.Find(&Authors).Error; err != nil {
			return Authors, err
		}
	}

	return Authors, nil
}

func AuthorCreation(author *models.Author) (*models.Author, error) {
	var existingAuthor models.Author
	err := DB.Where("name = ? AND email = ?", author.Name, author.Email).First(&existingAuthor).Error
	if err == nil {
		return nil, errors.New("author already exists")
	}
	if err := DB.Create(author).Error; err != nil {
		return nil, DB.Error
	}
	return author, nil
}

func UpdateAuthorInfo(updateAuthor *types.UpdateAuthorStruc) (*models.Author, error) {

	AuthorDetails := &models.Author{}
	AuthorDetails.ID = uint(updateAuthor.ID)
	if err := DB.Where("id = ?", AuthorDetails.ID).Find(AuthorDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no author found with given ID")
		}
	}

	if updateAuthor.Name != nil {
		AuthorDetails.Name = *updateAuthor.Name
	}
	if updateAuthor.Email != nil {
		AuthorDetails.Email = *updateAuthor.Email
	}
	if updateAuthor.Age != nil {
		AuthorDetails.Age = *updateAuthor.Age
	}

	err := DB.Save(AuthorDetails).Error
	if err != nil {
		return nil, err
	}
	return AuthorDetails, nil
}

func DeleteAuthor(ID int64) (string, error) {
	var author = models.Author{}
	DB.Unscoped().Where("id =?", ID).Delete(author)
	return "Delete successful", nil
}
