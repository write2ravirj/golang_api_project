package services

import (
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
)

// GenreService defines the methods required for a genre service.
type GenreService interface {
	GetGenres() ([]models.Genre, error)
}

type genreServiceImpl struct {
	Repo *repositories.GenreRepository
}

// NewGenreService initializes the GenreService
func NewGenreService(repo *repositories.GenreRepository) GenreService {
	return &genreServiceImpl{Repo: repo}
}

// GetGenres fetches the list of genres from the repository
func (s *genreServiceImpl) GetGenres() ([]models.Genre, error) {
	return s.Repo.GetGenres()
}
