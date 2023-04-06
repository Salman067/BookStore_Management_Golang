package routes

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthorStoreRoutes = func(router *mux.Router, authController *controllers.AuthorController) {
	router.HandleFunc("/author", authController.CreateAuthor).Methods("POST")
	router.HandleFunc("/author", authController.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{authorId}", authController.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{authorId}", authController.DeleteAuthor).Methods("DELETE")
}
