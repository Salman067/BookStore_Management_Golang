package routes

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router, bookControtter *controllers.BookController) {

	router.HandleFunc("/book", bookControtter.CreateBook).Methods("POST")
	router.HandleFunc("/book", bookControtter.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookControtter.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bookControtter.DeleteBook).Methods("DELETE")
}
