package domain

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type  IBookService  interface{
	GetBooks(reqStruct *types.BookReqStruc) ([]models.Book,error)
	CreateBook(book *types.CreateBookStruc) (*models.Book,error)
	UpdateBook(book *types.UpdateBookStruc) (*models.Book,error)
	DeleteBook(bookID string) (string, error)
}
type  IBookRepo interface{
	GetBooks(filterStruct *types.FilterBookStruc) ([]models.Book,error)
	CreateBook(book *types.CreateBookStruc) (*models.Book,error)
	UpdateBook(book *models.Book) (*models.Book,error)
	DeleteBook(bookID int64) (string, error)
}