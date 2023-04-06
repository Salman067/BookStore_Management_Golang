package models

import (
	"errors"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true" json:"author_id,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Email      string `gorm:"unique;not null" json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
}

func authorNameValidate(author string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 6 || len(name) > 200 {
			return errors.New("Please enter valid name")
		}
		return nil
	}
}

func authorEmailValidate(emai string) validation.RuleFunc {
	return func(value interface{}) error {
		email := value.(string)
		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(pattern)
		err := re.MatchString(email)
		if !err {
			return errors.New("Invalid email address")
		}
		return nil
	}
}

func (author Author) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.AuthorName, validation.Required,
			validation.By(authorNameValidate(author.AuthorName))),
		validation.Field(&author.Gender, validation.Required, validation.In("Female", "Male")),
		validation.Field(&author.Email, validation.By(authorEmailValidate(author.Email))),
	)
}
