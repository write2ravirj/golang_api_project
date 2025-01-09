package services

import (
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
)

// SizeService defines the methods required for a size service.
type SizeService interface {
	GetSizes() ([]models.Size, error)
}

type sizeServiceImpl struct {
	Repo *repositories.SizeRepository
}

// NewSizeService initializes the SizeService
func NewSizeService(repo *repositories.SizeRepository) SizeService {
	return &sizeServiceImpl{Repo: repo}
}

// GetSizes fetches the list of sizes from the repository
func (s *sizeServiceImpl) GetSizes() ([]models.Size, error) {
	return s.Repo.GetSizes()
}
