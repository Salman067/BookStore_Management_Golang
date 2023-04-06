package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/utils"
	"github.com/gorilla/mux"
)

//var BookService domain.BookServiceInterface

type BookController struct {
	bookService domain.BookServiceInterface
}

func SetBookService(bookService *domain.BookServiceInterface) *BookController {
	return &BookController{
		bookService: *bookService,
	}
}

func (bookController *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "pkglication/json")
	authorID := r.URL.Query().Get("author_id")
	ID := r.URL.Query().Get("id")
	bookID, bookError := strconv.Atoi(string(ID))
	bookAuthorID, authorError := strconv.Atoi(string(authorID))
	if bookError != nil && ID != "" {
		Response(w, 406, "Message : Invalid book id...!!")
		return
	} else if authorError != nil && authorID != "" {
		Response(w, 406, "Message : Invalid author id...!!")
		return
	}
	bookName := r.URL.Query().Get("book_name")
	bookPublication := r.URL.Query().Get("publication")
	books := types.ResponseBook{
		ID:          uint(bookID),
		AuthorID:    uint(bookAuthorID),
		BookName:    bookName,
		Publication: bookPublication,
	}
	newBook, dbError := bookController.bookService.GetBookService(&books)
	if dbError != nil {
		Response(w, 404, dbError.Error())
		return
	}
	res, err := json.Marshal(newBook)
	if err != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	if len(newBook) == 0 {
		Response(w, 404, "Message : Book not found...!!")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (bookController *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	newBook := models.Book{
		BookName:    CreateBook.BookName,
		AuthorID:    CreateBook.AuthorID,
		Publication: CreateBook.Publication,
	}
	err := newBook.Validate()
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	book, err := bookController.bookService.CreateBookService(&newBook)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	_, bookError := json.Marshal(book)
	if bookError != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	Response(w, 200, "Message : Book created successful..!!")
}

func (bookController *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, bookError := strconv.ParseInt(bookId, 0, 0)
	if bookError != nil {
		Response(w, 406, "Message : Invalid book ID")
		return
	}
	book, err := bookController.bookService.DeleteBookService(ID)
	if err != nil {
		Response(w, http.StatusBadRequest, "Message : Book id not found..!!")
		return
	}
	_, bookErr := json.Marshal(book)
	if bookErr != nil {
		Response(w, 406, "Marshalling error")
		return
	}
	Response(w, 200, "Message : Book deleted successful..!!")

}

func (bookController *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	json.NewDecoder(r.Body).Decode(&updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, bookError := strconv.ParseInt(bookId, 0, 0)
	if bookError != nil {
		Response(w, 404, "Message : Invalid book id..!!")
		return
	}
	fmt.Println(updateBook.AuthorID)
	er := updateBook.Validate()
	if er != nil {
		Response(w, 404, er.Error())
		return
	}
	_, err := bookController.bookService.UpdateBookService(*updateBook, ID)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	Response(w, 200, "Message : Book updated successful....!!")

}
