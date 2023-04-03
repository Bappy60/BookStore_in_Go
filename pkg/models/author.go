package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
	Books []Book `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

func (author Author) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&author.Email, validation.Required, is.Email),
		validation.Field(&author.Age, validation.Min(1)),
	)
}
