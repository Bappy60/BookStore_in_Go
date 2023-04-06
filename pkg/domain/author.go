package domain

import (
	"net/http"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)
type  IAuthorController interface{
	GetAuthors(w http.ResponseWriter, r *http.Request)
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
}

type  IAuthorService interface{
	GetAuthor(authorStruc *types.AuthorReqStruc) ([]models.Author, error)
	CreateAuthor(newAuthor *types.CreateAuthorStruc) (*models.Author, error)
	UpdateAuthor(updateAuthor *types.UpdateAuthorStruc) (*models.Author, error)
	DeleteAuthor(ID string) (string, error)
}

type  IAuthorRepo interface{
	GetAuthor(authorStruc *types.FilterAuthorStruc) ([]models.Author, error)
	CreateAuthor(newAuthor *types.CreateAuthorStruc) (*models.Author, error)
	UpdateAuthor(updateAuthor *models.Author) (*models.Author, error)
	DeleteAuthor(ID int64) (string, error)
}