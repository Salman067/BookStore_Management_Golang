package models

import (
	"fmt"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.Connect()
	fmt.Println("Database connected...")
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db = config.GetDB()
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {

	var Books []Book
	//fmt.Println("Hi")
	db = config.GetDB()
	db.Find(&Books)
	return Books

}

func GetBookById(Id int64) (*Book, *gorm.DB) {

	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db

}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
