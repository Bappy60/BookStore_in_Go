package models

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Author struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
	Books []Book `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

func (author Author) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.Name, validation.Required, validation.Length(3, 50),validation.By(AuthorNameValidator)),
		validation.Field(&author.Email, validation.Required, is.Email),
		validation.Field(&author.Age, validation.Min(1)),
	)
}
func AuthorNameValidator (value interface{}) error{
	if str, ok := value.(string); ok {
		match, err := regexp.MatchString("^[A-Za-z ]+$", str)
		if err != nil {
			return err
		}
		if !match {
			return  errors.New("author name should contain only letters and spaces")
		}
	}
	return nil
}