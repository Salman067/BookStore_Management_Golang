package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/utils"
	"github.com/gorilla/mux"
)

// var AuthorService domain.AuthorSeviceInterface

type AuthorController struct {
	authorService domain.AuthorSeviceInterface
}

func SetAuthorService(authorService *domain.AuthorSeviceInterface) *AuthorController {
	// AuthorService = *authorService
	return &AuthorController{
		authorService: *authorService,
	}
}

func (authController *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	createAuthor := &models.Author{}
	utils.ParseBody(r, createAuthor)
	newAuthor := models.Author{
		AuthorName: createAuthor.AuthorName,
		Gender:     createAuthor.Gender,
		Email:      createAuthor.Email,
		Address:    createAuthor.Address,
	}
	err := newAuthor.Validate()
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	author, err := authController.authorService.CreateAuthorService(&newAuthor)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	_, errAuthor := json.Marshal(author)
	if errAuthor != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	Response(w, 200, "Message : Author created successful..!!")

}

func (authController *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "pkglication/json")
	ID := r.URL.Query().Get("id")
	AuthorID, authorError := strconv.ParseInt(ID, 0, 0)
	if authorError != nil && ID != "" {
		Response(w, 406, "Message : Invalid author id...!!")
		return
	}
	AuthorName := r.URL.Query().Get("author_name")
	AuthorGender := r.URL.Query().Get("gender")
	AuthorEmail := r.URL.Query().Get("email")
	AuthorAddress := r.URL.Query().Get("address")
	getAuthor := types.ResponseAuthor{
		ID:         uint(AuthorID),
		AuthorName: AuthorName,
		Gender:     AuthorGender,
		Email:      AuthorEmail,
		Address:    AuthorAddress,
	}
	newAuthor, dbError := authController.authorService.GetAuthorService(&getAuthor)
	if dbError != nil {
		Response(w, 404, dbError.Error())
	}
	res, err := json.Marshal(newAuthor)
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	if len(newAuthor) == 0 {
		Response(w, 404, "Message : Author not found...!!")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (authController *AuthorController) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)

	if err != nil {
		Response(w, 406, "Message : Invalid author ID")
		return
	}
	author, err := authController.authorService.DeleteAuthorService(ID)
	if err != nil {
		Response(w, http.StatusBadRequest, "Message : Author id not found....!")
		return
	}
	_, authorErr := json.Marshal(author)
	if authorErr != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	Response(w, 200, "Message : Author deleted successful..!!")
}

func (authController *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	getAuthor := &models.Author{}
	json.NewDecoder(r.Body).Decode(&getAuthor)
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, authorError := strconv.ParseInt(authorId, 0, 0)
	if authorError != nil {
		Response(w, 404, "Message : Invalid author id..!!")
		return
	}
	if getAuthor.Email != "" {
		Response(w, 404, "Message : Email not updated..!!")
		return
	}
	// er := getAuthor.Validate()
	// if er != nil {
	// 	Response(w, 404, er.Error())
	// 	return
	// }
	_, err := authController.authorService.UpdateAuthorService(*getAuthor, ID)

	if err != nil {
		Response(w, http.StatusBadRequest, "Message : Author id not found....!")
		return
	}
	Response(w, 200, "Message : Author updated successful....!!")

}

func Response(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-type", "pkglication/json")
	res, err := json.Marshal(msg)
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	w.WriteHeader(statusCode)
	w.Write(res)
}
