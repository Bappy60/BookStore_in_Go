package models

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&b.Author, validation.Required, validation.Length(3, 50)),
		validation.Field(&b.Publication),
	)
}

func Initialize() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
