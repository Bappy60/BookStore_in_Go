package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ResponseStruc struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type FilterStruc struct {
	ID              uint
	Name            *string
	PublicationYear int
	NumberOfPages   int
	AuthorID        *uint
	Publication     *string
}
type AuthorStruc struct {
	ID    uint64 `gorm:"primaryKey" json:"author_id"`
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
}
type CreateAuthorStruc struct {
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
}

func (createAuthorStruc CreateAuthorStruc) Validate() error {
	return validation.ValidateStruct(&createAuthorStruc,
		validation.Field(&createAuthorStruc.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&createAuthorStruc.Email, validation.Required, is.Email),
		validation.Field(&createAuthorStruc.Age, validation.Min(1)),
	)
}

type UpdateBookStruc struct {
	ID              uint
	Name            *string `json:"name"`
	PublicationYear *int    `json:"publication_year"`
	NumberOfPages   *int    `json:"number_of_pages"`
	Publication     *string `json:"publication"`
}

func (updateBookStruc UpdateBookStruc) Validate() error {
	return validation.ValidateStruct(&updateBookStruc,
		validation.Field(&updateBookStruc.Name, validation.Length(3, 50)),
		validation.Field(&updateBookStruc.NumberOfPages, validation.Min(5)),
		validation.Field(&updateBookStruc.Publication, validation.Length(5, 50)),
		validation.Field(&updateBookStruc.PublicationYear, validation.Min(1200)),
	)
}

type UpdateAuthorStruc struct {
	ID    uint64
	Name  *string
	Email *string
	Age   *int
}

func (updateAuthorStruc UpdateAuthorStruc) Validate() error {
	return validation.ValidateStruct(&updateAuthorStruc,
		validation.Field(&updateAuthorStruc.Name, validation.Length(3, 50)),
		validation.Field(&updateAuthorStruc.Email, is.Email),
		validation.Field(&updateAuthorStruc.Age, validation.Min(5), validation.Max(150)),
	)
}

// func requiredIfNotZero(value interface{}) error {
// 	if intValue, ok := value.(int); ok && intValue == 0 {
// 		log.Println("entered in age validation")
// 		return errors.New("must not be zero")
// 	}
// 	return nil
// }
