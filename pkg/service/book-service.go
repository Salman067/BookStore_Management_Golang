package service

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
)

// var BookRepo domain.BookRepoInterface

// func SetBookRepo(bookRepo *domain.BookRepoInterface) {
// 	BookRepo = *bookRepo
// }

type BookService struct {
	repo domain.BookRepoInterface
}

func BookServiceInstance(bookRepo domain.BookRepoInterface) domain.BookServiceInterface {
	return &BookService{
		repo: bookRepo,
	}
}
func (service *BookService) CreateBookService(book *models.Book) (*types.ResponseBook, error) {
	bookService, err := service.repo.CreateBook(book)
	return bookService, err
}

func (service *BookService) GetBookService(resBook *types.ResponseBook) ([]types.Response, error) {
	books, err := service.repo.GetBooks(*resBook)
	return books, err
}

func (service *BookService) DeleteBookService(ID int64) (*types.ResponseBook, error) {
	book, err := service.repo.DeleteBook(ID)
	return book, err
}

func (service *BookService) UpdateBookService(updateBook models.Book, ID int64) (*models.Book, error) {
	bookUdate, err := service.repo.UpdateBook(&updateBook, ID)
	return bookUdate, err
}
