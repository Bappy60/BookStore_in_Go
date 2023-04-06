package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Bappy60/BookStore_in_Go/pkg/domain"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

type AuthorController struct {
	authorService domain.IAuthorService
}

func AuthorControllerInstance(authorService domain.IAuthorService) domain.IAuthorController {
	return &AuthorController{
		authorService: authorService,
	}
}

func (authorController *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorId := r.URL.Query().Get("author_id")
	authorName := r.URL.Query().Get("author_name")
	authorEmail := r.URL.Query().Get("email")
	authorAge := r.URL.Query().Get("author_age")

	authorReqstruc := types.AuthorReqStruc{
		ID:    authorId,
		Name:  authorName,
		Email: authorEmail,
		Age:   authorAge,
	}

	Authors, err := authorController.authorService.GetAuthor(&authorReqstruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(Authors)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	if len(Authors) == 0 {
		http.Error(w, "no author registered ", http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (authorController *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateAuthor := types.CreateAuthorStruc{}
	err := json.NewDecoder(r.Body).Decode(&CreateAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	err1 := CreateAuthor.Validate()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusNotAcceptable)
		return
	}

	author, err := authorController.authorService.CreateAuthor(&CreateAuthor)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	res, err := json.Marshal(author)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (authorController *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateAuthor = &types.UpdateAuthorStruc{}
	err := json.NewDecoder(r.Body).Decode(updateAuthor)
	if err != nil {
		http.Error(w, "Invalid Format of Data while decoding", http.StatusNotAcceptable)
		return
	}
	err1 := updateAuthor.Validate()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusNotAcceptable)
		return
	}
	vars := mux.Vars(r)
	authorId := vars["author_id"]
	updateAuthor.ID = authorId

	updatedAuthor, err := authorController.authorService.UpdateAuthor(updateAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedAuthor)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func (authorController *AuthorController) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	authorId := vars["author_id"]
	if authorId == "" {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}

	msg, err := authorController.authorService.DeleteAuthor(authorId)
	if err != nil {
		http.Error(w, "There is no author registered by this ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
}
