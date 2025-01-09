package services

import (
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
)

// AuthorService defines the methods required for an author service.
type AuthorService interface {
	GetAuthors() ([]models.Author, error)
}

type authorServiceImpl struct {
	Repo *repositories.AuthorRepository
}

func NewAuthorService(repo *repositories.AuthorRepository) AuthorService {
	return &authorServiceImpl{Repo: repo}
}

func (s *authorServiceImpl) GetAuthors() ([]models.Author, error) {
	return s.Repo.GetAuthors()
}
