package service

import (
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/domain"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/models"
	"github.com/MDABUSALMANHOSSAIN2018831067/BookStore_Management_Golang/pkg/types"
)

type AuthorService struct {
	repo domain.AuthorRepoInterface
}

func AuthorServiceInstance(authorRepo domain.AuthorRepoInterface) domain.AuthorSeviceInterface {
	return &AuthorService{
		repo: authorRepo,
	}
}

func (service *AuthorService) CreateAuthorService(author *models.Author) (*types.ResponseAuthor, error) {
	authorService, err := service.repo.CreateAuthor(author)
	return authorService, err
}

func (service *AuthorService) GetAuthorService(resAuthor *types.ResponseAuthor) ([]types.ResponseAuthor, error) {
	authors, err := service.repo.GetAuthors(*resAuthor)
	return authors, err
}

func (service *AuthorService) DeleteAuthorService(ID int64) (*types.ResponseAuthor, error) {
	author, err := service.repo.DeleteAuthor(ID)
	return author, err
}

func (service *AuthorService) UpdateAuthorService(updateAuthor models.Author, ID int64) (*models.Author, error) {
	authorUpdate, err := service.repo.UpdateAuthor(&updateAuthor, ID)
	return authorUpdate, err
}
