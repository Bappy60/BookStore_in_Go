package services

import (
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type BookService struct {
	repo domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		repo: bookRepo,
	}
}

func (service *BookService) GetBooks(reqStruct *types.BookReqStruc) ([]models.Book, error) {


	parsedId, err := strconv.ParseInt(reqStruct.ID, 0, 0)
	if err != nil && reqStruct.ID != "" {
		return nil,err
	}
	parsedAuthorID, err := strconv.ParseUint(reqStruct.AuthorID, 0, 0)
	if err != nil && reqStruct.AuthorID != "" {
		return nil,err
	}
	parsedNumOfPages, err := strconv.ParseInt(reqStruct.NumberOfPages, 0, 0)
	if err != nil && reqStruct.NumberOfPages != "" {
		return nil,err
	}
	parsedPublicationYear, err := strconv.ParseInt(reqStruct.PublicationYear, 0, 0)
	if err != nil && reqStruct.PublicationYear != "" {
		return nil,err
	}

	pAuthorID := uint(parsedAuthorID)
	fbookstruc := types.FilterBookStruc{
		ID:              uint(parsedId),
		Name:            &reqStruct.Name,
		PublicationYear: int(parsedPublicationYear),
		NumberOfPages:   int(parsedNumOfPages),
		AuthorID:        &pAuthorID,
		Publication:     &reqStruct.Publication,
	}

	res, err := service.repo.GetBooks(&fbookstruc)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *BookService) CreateBook(book *types.CreateBookStruc) (*models.Book, error) {

	res, err := service.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *BookService) UpdateBook(reqBook *types.UpdateBookStruc) (*models.Book, error) {
	bookID, err := strconv.ParseUint(reqBook.ID, 0, 0)
	if err != nil {
		return nil,err
	}
	book := &models.Book{
		ID: uint(bookID),
		Name: reqBook.Name,
		NumberOfPages: reqBook.NumberOfPages,
		Publication: reqBook.Publication,
		PublicationYear: reqBook.PublicationYear,
	}
	res, err := service.repo.UpdateBook(book)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *BookService) DeleteBook(bookId string) (string, error) {

	parsedId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return "invalid format of ID",err
	}
	res, err := service.repo.DeleteBook(parsedId)
	if err != nil {
		return "", err
	}
	return res, nil
}
