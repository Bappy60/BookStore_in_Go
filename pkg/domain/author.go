package domain

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type  IAuthorRepo interface{
	GetAuthor(authorStruc *types.AuthorStruc) ([]models.Author, error)
	AuthorCreation(newAuthor *types.CreateAuthorStruc) (*models.Author, error)
	UpdateAuthorInfo(updateAuthor *types.UpdateAuthorStruc) (*models.Author, error)
	DeleteAuthor(ID int64) (string, error)
}