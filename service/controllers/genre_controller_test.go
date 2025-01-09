package controllers

import (
	"encoding/json"
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockGenreService is a mock implementation of the GenreService
type MockGenreService struct {
	mock.Mock
}

func (m *MockGenreService) GetGenres() ([]models.Genre, error) {
	args := m.Called()

	// Check if args.Get(0) is nil or not of the expected type
	if genres, ok := args.Get(0).([]models.Genre); ok {
		return genres, args.Error(1)
	}

	return nil, args.Error(1) // Return nil for genres if type assertion fails
}

func TestGenreController_GetGenres(t *testing.T) {
	mockGenres := []models.Genre{
		{ID: 1, Title: "Fiction"},
		{ID: 2, Title: "Non-Fiction"},
		{ID: 3, Title: "Science Fiction"},
	}

	t.Run("successfully fetches genres", func(t *testing.T) {
		// Arrange
		mockService := new(MockGenreService)
		mockService.On("GetGenres").Return(mockGenres, nil)

		controller := NewGenreController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/genres", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetGenres(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)

		var genres []models.Genre
		err := json.Unmarshal(rec.Body.Bytes(), &genres)
		assert.NoError(t, err)
		assert.Equal(t, mockGenres, genres)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockGenreService)
		mockService.On("GetGenres").Return(nil, errors.New("service error"))

		controller := NewGenreController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/genres", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetGenres(rec, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Error fetching genres")

		mockService.AssertExpectations(t)
	})
}
