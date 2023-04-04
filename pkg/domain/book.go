package domain

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type  IBookRepo interface{
	GetBooks(filterStruct *types.FilterStruc) ([]models.Book,error)
	CreateBook(book *types.CreateBookStruc) (*models.Book,error)
	UpdateBook(book *types.UpdateBookStruc) (*models.Book,error)
	DeleteBook(bookID int64) (string, error)
}