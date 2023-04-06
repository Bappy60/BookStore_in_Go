package types

import (
	"errors"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BookReqStruc struct {
	ID              string
	Name            string
	PublicationYear string
	NumberOfPages   string
	AuthorID        string
	Publication     string
}

type AuthorReqStruc struct {
	ID    string
	Name  string
	Email string
	Age   string
}

type FilterBookStruc struct {
	ID              uint
	Name            *string
	PublicationYear int
	NumberOfPages   int
	AuthorID        *uint
	Publication     *string
}

type FilterAuthorStruc struct {
	ID    uint
	Name  string
	Email string
	Age   int64
}

type CreateBookStruc struct {
	Name            string `gorm:"" json:"name"`
	PublicationYear int    `json:"publication_year"`
	NumberOfPages   int    `json:"number_of_pages"`
	AuthorID        uint   `gorm:"index" json:"author_id"` // foreign key
	Publication     string `json:"publication"`
}

func (book CreateBookStruc) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&book.PublicationYear, validation.Required),
		validation.Field(&book.NumberOfPages, validation.Required),
		validation.Field(&book.AuthorID, validation.Required),
		validation.Field(&book.Publication, validation.Length(1, 50)),
	)
}

type CreateAuthorStruc struct {
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
}

func (createAuthorStruc CreateAuthorStruc) Validate() error {
	return validation.ValidateStruct(&createAuthorStruc,
		validation.Field(&createAuthorStruc.Name, validation.When(createAuthorStruc.Name != "", validation.By(AuthorNameValidator)), validation.Length(3, 50)),
		validation.Field(&createAuthorStruc.Email, is.Email),
		validation.Field(&createAuthorStruc.Age, validation.Min(int(0)).Exclusive(), validation.NilOrNotEmpty.Error("must be greater than 0"), validation.Max(150)))
	
}

type UpdateBookStruc struct {
	ID              string
	Name            string `json:"name"`
	PublicationYear int    `json:"publication_year"`
	NumberOfPages   int    `json:"number_of_pages"`
	Publication     string `json:"publication"`
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
	ID    string
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
}

func (updateAuthorStruc UpdateAuthorStruc) Validate() error {
	return validation.ValidateStruct(&updateAuthorStruc,
		validation.Field(&updateAuthorStruc.Name, validation.When(updateAuthorStruc.Name != "", validation.By(AuthorNameValidator)), validation.Length(3, 50)),
		validation.Field(&updateAuthorStruc.Email, is.Email),
		validation.Field(&updateAuthorStruc.Age, validation.Min(int(0)).Exclusive(), validation.NilOrNotEmpty.Error("must be greater than 0"), validation.Max(150)))
}

func AuthorNameValidator(value interface{}) error {
	if str, ok := value.(string); ok {
		match, err := regexp.MatchString("^[A-Za-z ]+$", str)
		if err != nil {
			return err
		}
		if !match {
			return errors.New("author name should contain only letters and spaces")
		}
	}
	return nil
}
