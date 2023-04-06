package types

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
)

type Response struct {
	models.Book
	models.Author
}

type AuthorResponseStruct struct {
	// ID         uint   `json:"author_id,omitempty"`
	// AuthorName string `json:"author_name,omitempty"`
	// Gender     string `json:"author_gender,omitempty"`
	// Email      string `json:"author_email,omitempty"`
	// Address    string `json:"author_address,omitempty"`
	models.Author
}

type ResponseBook struct {
	ID          uint   `json:"book_id,omitempty"`
	BookName    string `json:"book_name,omitempty"`
	AuthorID    uint   `json:"author_id,omitempty"`
	Publication string `json:"publication,omitempty"`
}

type ResponseAuthor struct {
	ID         uint   `json:"author_id,omitempty"`
	AuthorName string `json:"author_name ,omitempty" `
	Gender     string `json:"gender,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"aadress,omitempty"`
}
