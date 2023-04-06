package repository

import (
	"errors"
	"fmt"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

func BookDBInterface(db *gorm.DB) domain.BookRepoInterface {
	return &BookRepo{
		DB: db,
	}

}

func (bookRepo *BookRepo) CreateBook(book *models.Book) (*types.ResponseBook, error) {
	author := models.Author{}
	err := bookRepo.DB.Table("authors").
		Where("id=?", book.AuthorID).First(&author).Error
	if err != nil {
		return nil, errors.New("Message : Author id not found..!!")
	}
	if err := bookRepo.DB.Where("books.book_name = ? AND books.author_id = ?",
		book.BookName, book.AuthorID).
		First(&book).Error; err == nil {
		return nil, errors.New("Message : Book already exists..!")
	}
	if err := bookRepo.DB.Table("books").Create(&book).Error; err != nil {
		return nil, err
	}
	responseBook := types.ResponseBook{
		ID:          book.ID,
		BookName:    book.BookName,
		AuthorID:    book.AuthorID,
		Publication: book.Publication,
	}
	return &responseBook, nil

}

func (bookRepo *BookRepo) GetBooks(resBook types.ResponseBook) ([]types.Response, error) {
	var getBook []types.Response
	if resBook.ID != 0 || resBook.AuthorID != 0 || resBook.BookName != "" || resBook.Publication != "" {
		if resBook.BookName != "" {
			bookRepo.DB.Table("books").Select("*").
				Joins("join authors on books.author_id = authors.id").
				Where("books.book_name = ?", resBook.BookName).Find(&getBook)
		}

		if resBook.AuthorID != 0 {
			bookRepo.DB.Table("books").Select("*").
				Joins("join authors on books.author_id = authors.id").
				Where("books.author_id = ?", resBook.AuthorID).Find(&getBook)
		}

		if resBook.ID != 0 {
			bookRepo.DB.Table("books").Select("*").
				Joins("join authors on books.author_id = authors.id").
				Where("books.id = ?", resBook.ID).Find(&getBook)
		}

		if resBook.Publication != "" {
			bookRepo.DB.Table("books").Select("*").
				Joins("join authors on books.author_id = authors.id").
				Where("books.publication = ?", resBook.Publication).Find(&getBook)
		}
		return getBook, nil
	}

	if err := bookRepo.DB.Table("books").Select("*").
		Joins("join authors on books.author_id = authors.id").Find(&getBook).Error; err != nil {
		return nil, err
	}
	return getBook, nil
}

func (bookRepo *BookRepo) UpdateBook(book *models.Book, ID int64) (*models.Book, error) {
	var getBook models.Book
	var getAuthor models.Author
	fmt.Println(book.AuthorID)
	if err := bookRepo.DB.Table("books").
		Where("books.id = ? ", ID).Find(&getBook).Error; err != nil {
		return nil, errors.New("Message : Book id not found..!!")
	}
	if err := bookRepo.DB.Table("authors").
		Where("authors.id = ? ", book.AuthorID).Find(&getAuthor).Error; err != nil {
		return nil, errors.New("Message : Author id not found..!!")
	}
	if book.BookName != "" || book.Publication != "" ||
		book.AuthorID != 0 || book.ID != 0 {
		if getBook.BookName != "" {
			getBook.BookName = book.BookName
		}
		if getBook.AuthorID != 0 {
			getBook.AuthorID = book.AuthorID
		}
		if getBook.Publication != "" {
			getBook.Publication = book.Publication
		}
	}
	bookRepo.DB.Save(&getBook)
	fmt.Println(getBook)
	return &getBook, nil
}

func (bookRepo *BookRepo) DeleteBook(ID int64) (*types.ResponseBook, error) {
	var book models.Book
	bookRes := types.ResponseBook{
		ID: book.ID,
	}
	if err := bookRepo.DB.Table("books").Where("id = ?", ID).Find(&book).Error; err != nil {
		return nil, err
	}

	if err := bookRepo.DB.Unscoped().Where("id = ?", ID).Delete(book).Error; err != nil {
		return nil, err
	}
	return &bookRes, nil
}
