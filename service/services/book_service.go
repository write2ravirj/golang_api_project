package services

import (
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
)

// BookService defines the methods required for a book service.
type BookService interface {
	GetBooks(filters map[string]interface{}) ([]models.Book, error)
}

type bookServiceImpl struct {
	Repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) BookService {
	return &bookServiceImpl{Repo: repo}
}

func (s *bookServiceImpl) GetBooks(filters map[string]interface{}) ([]models.Book, error) {
	return s.Repo.GetBooks(filters)
}
