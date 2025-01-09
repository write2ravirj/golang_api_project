package controllers

import (
	"encoding/json"
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthorService is a mock implementation of the AuthorService
type MockAuthorService struct {
	mock.Mock
}

func (m *MockAuthorService) GetAuthors() ([]models.Author, error) {
	args := m.Called()

	// Check if args.Get(0) is nil or not of the expected type
	if authors, ok := args.Get(0).([]models.Author); ok {
		return authors, args.Error(1)
	}

	return nil, args.Error(1) // Return nil for authors if type assertion fails
}

func TestAuthorController_GetAuthors(t *testing.T) {
	// Mock data
	mockAuthors := []models.Author{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	t.Run("successfully fetches authors", func(t *testing.T) {
		// Arrange
		mockService := new(MockAuthorService)
		mockService.On("GetAuthors").Return(mockAuthors, nil)

		controller := NewAuthorController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/authors", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetAuthors(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)

		var authors []models.Author
		err := json.Unmarshal(rec.Body.Bytes(), &authors)
		assert.NoError(t, err)
		assert.Equal(t, mockAuthors, authors)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockAuthorService)
		mockService.On("GetAuthors").Return(nil, errors.New("service error"))

		controller := NewAuthorController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/authors", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetAuthors(rec, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Error fetching authors")

		mockService.AssertExpectations(t)
	})
}
