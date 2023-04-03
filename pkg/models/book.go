package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name            string `gorm:"" json:"name"`
	PublicationYear int    `json:"publication_year"`
	NumberOfPages   int    `json:"number_of_pages"`
	AuthorID        uint   `gorm:"index" json:"author_id"` // foreign key
	Author          Author `gorm:"foreignKey:AuthorID"`
	Publication     string `json:"publication"`
}




func (book Book) Validate() error {
	return validation.ValidateStruct(&book,
        validation.Field(&book.Name, validation.Required, validation.Length(3, 50)),
        validation.Field(&book.PublicationYear, validation.Required),
        validation.Field(&book.NumberOfPages, validation.Required),
        validation.Field(&book.AuthorID, validation.Required),
        validation.Field(&book.Publication, validation.Length(1, 50)),
	)
}
