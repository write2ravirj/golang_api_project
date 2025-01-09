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

// MockBookService is a mock implementation of the BookService
type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetBooks(filters map[string]interface{}) ([]models.Book, error) {
	args := m.Called(filters)

	// Extract arguments and return
	if authors, ok := args.Get(0).([]models.Book); ok {
		return authors, args.Error(1)
	}

	return nil, args.Error(1) // Return nil for books if type assertion fails
}

func TestBookController_GetBooks(t *testing.T) {
	mockBooks := []models.Book{
		{ID: 1, Title: "Book One", YearPublished: 2020, Rating: 4.5, Pages: 300},
		{ID: 2, Title: "Book Two", YearPublished: 2021, Rating: 4.0, Pages: 250},
	}

	t.Run("successfully fetches books", func(t *testing.T) {
		// Arrange
		mockService := new(MockBookService)
		mockService.On("GetBooks", mock.Anything).Return(mockBooks, nil)

		controller := NewBookController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/books?authors=1,2&min-pages=200&max-pages=400", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetBooks(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)

		var books []models.Book
		err := json.Unmarshal(rec.Body.Bytes(), &books)
		assert.NoError(t, err)
		assert.Equal(t, mockBooks, books)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockBookService)
		mockService.On("GetBooks", mock.Anything).Return(nil, errors.New("service error"))

		controller := NewBookController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/books?min-pages=100", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetBooks(rec, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Error fetching books")

		mockService.AssertExpectations(t)
	})

	t.Run("handles invalid query parameters gracefully", func(t *testing.T) {
		// Arrange
		mockService := new(MockBookService)
		mockService.On("GetBooks", mock.Anything).Return(mockBooks, nil)

		controller := NewBookController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/books?min-pages=invalid", nil)
		rec := httptest.NewRecorder()

		// Act
		controller.GetBooks(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code) // Invalid params are ignored, so it defaults to fetching
		mockService.AssertExpectations(t)
	})
}
