package models

import (
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"book_id,omitempty"`
	BookName    string `json:"book_name,omitempty"`
	Publication string `json:"publication"`
	AuthorID    uint   `json:"author_id,omitempty" gorm:"foreignKey:AuthorID ;references:ID;"`
}

func authorIDValidate(author uint) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(uint) == 0 {
			return errors.New("Enter valid author id(Numerical)")
		}
		return nil
	}
}

func bookNameValidate(bookName string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 4 || len(name) > 200 {
			return errors.New("Please enter valid book name")
		}
		return nil
	}
}

func pubNameValidate(pubName string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 4 || len(name) > 200 {
			return errors.New("Please enter valid publication")
		}
		return nil
	}
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.BookName, validation.Required, validation.Length(1, 200),
			validation.By(bookNameValidate(b.BookName)),
			validation.Required),
		validation.Field(&b.Publication,
			validation.Length(4, 150),
			validation.By(pubNameValidate(b.Publication)),
			validation.Required),
		validation.Field(&b.AuthorID,
			validation.By(authorIDValidate(b.AuthorID)),
			validation.Required),
	)
}
