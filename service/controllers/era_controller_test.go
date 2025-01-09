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

// MockErasService is a mock implementation of the ErasService
type MockErasService struct {
	mock.Mock
}

func (m *MockErasService) GetEras() ([]models.Era, error) {
	args := m.Called()

	// Check if args.Get(0) is nil or not of the expected type
	if eras, ok := args.Get(0).([]models.Era); ok {
		return eras, args.Error(1)
	}

	return nil, args.Error(1) // Return nil for eras if type assertion fails
}

func TestErasController_GetEras(t *testing.T) {
	minYear := 1800
	maxYear := 2000
	mockEras := []models.Era{
		{ID: 1, Title: "Ancient", MinYear: nil, MaxYear: nil},
		{ID: 2, Title: "Medieval", MinYear: &minYear, MaxYear: nil},
		{ID: 3, Title: "Modern", MinYear: &minYear, MaxYear: &maxYear},
	}

	t.Run("successfully fetches eras", func(t *testing.T) {
		// Arrange
		mockService := new(MockErasService)
		mockService.On("GetEras").Return(mockEras, nil)

		controller := NewErasController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/eras", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetEras(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)

		var eras []models.Era
		err := json.Unmarshal(rec.Body.Bytes(), &eras)
		assert.NoError(t, err)
		assert.Equal(t, mockEras, eras)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockErasService)
		mockService.On("GetEras").Return(nil, errors.New("service error"))

		controller := NewErasController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/eras", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetEras(rec, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Error fetching eras")

		mockService.AssertExpectations(t)
	})
}
