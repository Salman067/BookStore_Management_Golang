package domain

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
)

type BookRepoInterface interface {
	GetBooks(resBook types.ResponseBook) ([]types.Response, error)
	CreateBook(book *models.Book) (*types.ResponseBook, error)
	UpdateBook(b *models.Book, ID int64) (*models.Book, error)
	DeleteBook(ID int64) (*types.ResponseBook, error)
}
type BookServiceInterface interface {
	CreateBookService(book *models.Book) (*types.ResponseBook, error)
	GetBookService(resBook *types.ResponseBook) ([]types.Response, error)
	DeleteBookService(ID int64) (*types.ResponseBook, error)
	UpdateBookService(updateBook models.Book, ID int64) (*models.Book, error)
}

type AuthorRepoInterface interface {
	CreateAuthor(author *models.Author) (*types.ResponseAuthor, error)
	GetAuthors(FA types.ResponseAuthor) ([]types.ResponseAuthor, error)
	UpdateAuthor(updateAuthor *models.Author, ID int64) (*models.Author, error)
	DeleteAuthor(ID int64) (*types.ResponseAuthor, error)
}

type AuthorSeviceInterface interface {
	CreateAuthorService(author *models.Author) (*types.ResponseAuthor, error)
	GetAuthorService(resAuthor *types.ResponseAuthor) ([]types.ResponseAuthor, error)
	DeleteAuthorService(ID int64) (*types.ResponseAuthor, error)
	UpdateAuthorService(updateAuthor models.Author, ID int64) (*models.Author, error)
}
