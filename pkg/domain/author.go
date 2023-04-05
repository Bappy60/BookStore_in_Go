package domain

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type  IAuthorService interface{
	GetAuthor(authorStruc *types.AuthorReqStruc) ([]models.Author, error)
	AuthorCreation(newAuthor *types.CreateAuthorStruc) (*models.Author, error)
	UpdateAuthorInfo(updateAuthor *types.UpdateAuthorStruc) (*models.Author, error)
	DeleteAuthor(ID string) (string, error)
}

type  IAuthorRepo interface{
	GetAuthor(authorStruc *types.FilterAuthorStruc) ([]models.Author, error)
	AuthorCreation(newAuthor *types.CreateAuthorStruc) (*models.Author, error)
	UpdateAuthorInfo(updateAuthor *models.Author) (*models.Author, error)
	DeleteAuthor(ID int64) (string, error)
}