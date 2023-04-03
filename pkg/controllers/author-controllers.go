package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/Bappy60/BookStore_in_Go/pkg/repositories"
	"github.com/Bappy60/BookStore_in_Go/pkg/types"
	"github.com/gorilla/mux"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorId := r.URL.Query().Get("author_id")
	authorName := r.URL.Query().Get("author_name")
	authorEmail := r.URL.Query().Get("email")
	authorAge := r.URL.Query().Get("author_age")

	parsedId, err := strconv.ParseUint(authorId, 0, 0)
	if err != nil && authorId != "" {
		http.Error(w, "invalid format of data while parsing id", http.StatusBadRequest)
		return
	}
	parsedAge, err := strconv.ParseInt(authorAge, 0, 0)
	if err != nil && authorAge != "" {
		http.Error(w, "invalid format of data while parsing age", http.StatusBadRequest)
		return
	}

	Authorstruc := types.AuthorStruc{
		ID:    parsedId,
		Name:  authorName,
		Email: authorEmail,
		Age:   int(parsedAge),
	}

	Authors, err := repositories.GetAuthor(&Authorstruc)
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

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
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
	authorStruct := &models.Author{
		Name:  CreateAuthor.Name,
		Email: CreateAuthor.Email,
		Age:   CreateAuthor.Age,
	}
	author, err := repositories.AuthorCreation(authorStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(author)
	if err != nil {
		http.Error(w, "Error While Marshaling", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
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
	parsedAuthorId, err := strconv.ParseUint(authorId, 0, 0)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	updateAuthor.ID = parsedAuthorId

	updatedAuthor, err := repositories.UpdateAuthorInfo(updateAuthor)
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

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	authorId := vars["author_id"]
	if authorId == "" {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	parsedAuthorId, err := strconv.ParseUint(authorId, 0, 0)
	if err != nil {
		http.Error(w, "invalid format of ID", http.StatusBadRequest)
		return
	}
	AuthorStruc := types.AuthorStruc{
		ID: parsedAuthorId,
	}
	author, err := repositories.GetAuthor(&AuthorStruc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(author) == 0 {
		http.Error(w, "there is no author registered by this ID", http.StatusBadRequest)
		return
	}
	msg, err := repositories.DeleteAuthor(int64(parsedAuthorId))
	if err != nil {
		http.Error(w, "There is no author registered by this ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
}
