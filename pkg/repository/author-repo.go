package repository

import (
	"errors"

	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
	"github.com/jinzhu/gorm"
)

type AuthorRepo struct {
	DB *gorm.DB
}

func AuthorDBInterface(db *gorm.DB) domain.AuthorRepoInterface {
	return &AuthorRepo{
		DB: db,
	}

}
func (authorRepo *AuthorRepo) CreateAuthor(author *models.Author) (*types.ResponseAuthor, error) {
	if err := authorRepo.DB.Where("email = ? AND author_name = ?",
		author.Email, author.AuthorName).First(&author).Error; err == nil {
		return nil, errors.New("Message : Author already exists..!")
	}
	if err := authorRepo.DB.Create(&author).Error; err != nil {
		return nil, err
	}
	responseAuthor := types.ResponseAuthor{
		ID:         author.ID,
		AuthorName: author.AuthorName,
		Gender:     author.Gender,
		Email:      author.Email,
		Address:    author.Address,
	}
	return &responseAuthor, nil
}

func (authorRepo *AuthorRepo) GetAuthors(FA types.ResponseAuthor) ([]types.ResponseAuthor, error) {
	var getAuthor []types.ResponseAuthor
	if FA.ID != 0 || FA.AuthorName != "" || FA.Gender != "" ||
		FA.Email != "" || FA.Address != "" {
		if FA.ID != 0 {
			authorRepo.DB.Table("authors").Select("*").
				Where("authors.id = ?", FA.ID).Find(&getAuthor)
		}
		if FA.AuthorName != "" {
			authorRepo.DB.Table("authors").Select("*").
				Where("authors.author_name = ?", FA.AuthorName).Find(&getAuthor)
		}
		if FA.Gender != "" {
			authorRepo.DB.Table("authors").Select("*").
				Where("authors.gender = ?", FA.Gender).Find(&getAuthor)
		}
		if FA.Email != "" {
			authorRepo.DB.Table("authors").Select("*").
				Where("authors.email = ?", FA.Email).Find(&getAuthor)
		}
		if FA.Address != "" {
			authorRepo.DB.Table("authors").Select("*").
				Where("authors.address = ?", FA.Address).Find(&getAuthor)
		}
		return getAuthor, nil
	}

	if err := authorRepo.DB.Table("authors").Select("*").
		Find(&getAuthor).Error; err != nil {
		return nil, err
	}
	return getAuthor, nil
}

func (authorRepo *AuthorRepo) DeleteAuthor(ID int64) (*types.ResponseAuthor, error) {
	var author models.Author
	authorRes := types.ResponseAuthor{
		ID: author.ID,
	}
	if err := authorRepo.DB.Table("authors").Where("authors.id = ?", ID).Find(&author).Error; err != nil {
		return nil, err
	}
	if err := authorRepo.DB.Unscoped().Where("authors.id = ?", ID).Delete(author).Error; err != nil {
		return nil, err
	}
	return &authorRes, nil
}

func (authorRepo *AuthorRepo) UpdateAuthor(updateAuthor *models.Author, ID int64) (*models.Author, error) {
	var author models.Author
	if err := authorRepo.DB.Where("authors.id = ? ", ID).Find(&author).Error; err != nil {
		return nil, err
	}

	if updateAuthor.AuthorName != "" {
		author.AuthorName = updateAuthor.AuthorName
	}
	if updateAuthor.Gender != "" {
		author.Gender = updateAuthor.Gender
	}
	if updateAuthor.Address != "" {
		author.Address = updateAuthor.Address
	}
	if err := authorRepo.DB.Save(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}
