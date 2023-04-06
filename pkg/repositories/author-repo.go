package repositories

import (
	"errors"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/jinzhu/gorm"
)

type authorRepo struct {
	db *gorm.DB
}

func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

func (repo *authorRepo) GetAuthor(authorStruc *types.FilterAuthorStruc) ([]models.Author, error) {
	var Authors []models.Author

	query := repo.db.Model(&models.Author{}).Preload("Books")
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

func (repo *authorRepo) CreateAuthor(newAuthor *types.CreateAuthorStruc) (*models.Author, error) {
	var author models.Author
	err := repo.db.Where("name = ? AND email = ?", newAuthor.Name, newAuthor.Email).First(&author).Error
	if err == nil {
		return nil, errors.New("author already exists")
	}

	author.Name = newAuthor.Name
	author.Email = newAuthor.Email
	author.Age = newAuthor.Age

	if err := repo.db.Create(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (repo *authorRepo) UpdateAuthor(updateAuthor *models.Author) (*models.Author, error) {

	AuthorDetails := &models.Author{}
	AuthorDetails.ID = updateAuthor.ID
	if err := repo.db.Where("id = ?", AuthorDetails.ID).Find(AuthorDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no author found with given ID")
		}
	}
	if updateAuthor.Name != "" {
		AuthorDetails.Name = updateAuthor.Name
	}
	if updateAuthor.Email != "" {
		AuthorDetails.Email = updateAuthor.Email
	}
	if updateAuthor.Age != 0 {
		AuthorDetails.Age = updateAuthor.Age
	}

	err := repo.db.Save(AuthorDetails).Error
	if err != nil {
		return nil, err
	}
	return AuthorDetails, nil
}

func (repo *authorRepo)  DeleteAuthor(ID int64) (string, error) {
	var author = models.Author{}
	if err := repo.db.Where("id = ?", ID).Find(&author).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("no author found with given ID")
		}
	}
	repo.db.Unscoped().Where("id =?", ID).Delete(author)
	return "Delete successful", nil
}
