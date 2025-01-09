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

// MockSizeService is a mock implementation of the SizeService
type MockSizeService struct {
	mock.Mock
}

func (m *MockSizeService) GetSizes() ([]models.Size, error) {
	args := m.Called()

	// Check if args.Get(0) is nil or not of the expected type
	if sizes, ok := args.Get(0).([]models.Size); ok {
		return sizes, args.Error(1)
	}

	return nil, args.Error(1) // Return nil for sizes if type assertion fails
}

func TestSizeController_GetSizes(t *testing.T) {
	// Example sizes, some with MinPages and MaxPages set, others without
	minPages := 100
	maxPages := 500
	mockSizes := []models.Size{
		{ID: 1, Title: "Small", MinPages: &minPages},
		{ID: 2, Title: "Medium", MinPages: nil, MaxPages: &maxPages},
		{ID: 3, Title: "Large", MinPages: nil, MaxPages: nil},
	}

	t.Run("successfully fetches sizes", func(t *testing.T) {
		// Arrange
		mockService := new(MockSizeService)
		mockService.On("GetSizes").Return(mockSizes, nil)

		controller := NewSizeController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/sizes", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetSizes(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)

		var sizes []models.Size
		err := json.Unmarshal(rec.Body.Bytes(), &sizes)
		assert.NoError(t, err)
		assert.Equal(t, mockSizes, sizes)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockSizeService)
		mockService.On("GetSizes").Return(nil, errors.New("service error"))

		controller := NewSizeController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/sizes", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetSizes(rec, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Error fetching sizes")

		mockService.AssertExpectations(t)
	})
}
