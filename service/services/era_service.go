package services

import (
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
)

// ErasService defines the methods required for an eras service.
type ErasService interface {
	GetEras() ([]models.Era, error)
}

type erasServiceImpl struct {
	Repo *repositories.ErasRepository
}

// NewErasService initializes the ErasService
func NewErasService(repo *repositories.ErasRepository) ErasService {
	return &erasServiceImpl{Repo: repo}
}

// GetEras fetches the list of eras from the repository
func (s *erasServiceImpl) GetEras() ([]models.Era, error) {
	return s.Repo.GetEras()
}
