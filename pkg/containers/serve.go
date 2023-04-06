package containers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/config"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/controllers"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/repository"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/routes"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/service"
	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	db := config.Connect()
	bookRepo := repository.BookDBInterface(db)
	bookService := service.BookServiceInstance(bookRepo)
	bookController := controllers.SetBookService(&bookService)

	authorRepo := repository.AuthorDBInterface(db)
	authorService := service.AuthorServiceInstance(authorRepo)
	authController := controllers.SetAuthorService(&authorService)

	// db.DropTableIfExists(&models.Author{})
	// db.DropTableIfExists(&models.Book{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})

	routes.RegisterBookStoreRoutes(r, bookController)
	routes.RegisterAuthorStoreRoutes(r, authController)
	http.Handle("/", r)
	fmt.Println("Server starting....")
	fmt.Println("Database connected...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
