package services

import (
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
)

type AuthorService struct {
	repo domain.IAuthorRepo
}

func AuthorServiceInstance(authorRepo domain.IAuthorRepo) domain.IAuthorService {
	return &AuthorService{
		repo: authorRepo,
	}
}


func (service *AuthorService) GetAuthor(reqStruc *types.AuthorReqStruc) ([]models.Author, error){

	parsedId, err := strconv.ParseUint(reqStruc.ID, 0, 0)
	if err != nil && reqStruc.ID != "" {
		//http.Error(w, "invalid format of data while parsing id", http.StatusBadRequest)
		return nil, err
	}
	parsedAge, err := strconv.ParseInt(reqStruc.Age, 0, 0)
	if err != nil && reqStruc.Age != "" {
		//http.Error(w, "invalid format of data while parsing age", http.StatusBadRequest)
		return nil, err
	}

	fauthorStruc := &types.FilterAuthorStruc{
		ID: uint(parsedId),
		Name: reqStruc.Name,
		Email: reqStruc.Email,
		Age: parsedAge,
	}
	res, err := service.repo.GetAuthor(fauthorStruc)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (service *AuthorService) AuthorCreation(newAuthor *types.CreateAuthorStruc) (*models.Author, error){

	res, err := service.repo.AuthorCreation(newAuthor)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (service *AuthorService) UpdateAuthorInfo(updateAuthor *types.UpdateAuthorStruc) (*models.Author, error){
	parsedAuthorId, err := strconv.ParseUint(updateAuthor.ID, 0, 0)
	if err != nil {
		//http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return nil,err
	}

	author := &models.Author{
		ID: uint(parsedAuthorId),
		Name: updateAuthor.Name,
		Email: updateAuthor.Email,
		Age: updateAuthor.Age,
	}

	res, err := service.repo.UpdateAuthorInfo(author)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (service *AuthorService) DeleteAuthor(authorId string) (string, error){
	parsedAuthorId, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		//http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return "",err
	}
	res, err := service.repo.DeleteAuthor(parsedAuthorId)
	if err != nil {
		return "", err
	}
	return res, nil
}