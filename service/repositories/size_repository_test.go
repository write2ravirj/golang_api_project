package repositories

import (
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeRepository_GetSizes(t *testing.T) {
	t.Run("successfully fetches sizes", func(t *testing.T) {
		// Mock database and data
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "title", "min_pages", "max_pages"}).
			AddRow(1, "Short", 1, 100).
			AddRow(2, "Medium", 101, 300).
			AddRow(3, "Long", 301, nil)

		mock.ExpectQuery("SELECT id, title, min_pages, max_pages FROM size").WillReturnRows(rows)

		// Initialize repository
		repo := NewSizeRepository(db)

		// Call GetSizes
		sizes, err := repo.GetSizes()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, sizes, 3)

		expected := []models.Size{
			{ID: 1, Title: "Short", MinPages: intPtr(1), MaxPages: intPtr(100)},
			{ID: 2, Title: "Medium", MinPages: intPtr(101), MaxPages: intPtr(300)},
			{ID: 3, Title: "Long", MinPages: intPtr(301), MaxPages: nil},
		}
		assert.Equal(t, expected, sizes)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles empty result set", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title, min_pages, max_pages FROM size").WillReturnRows(sqlmock.NewRows(nil))

		// Initialize repository
		repo := NewSizeRepository(db)

		// Call GetSizes
		sizes, err := repo.GetSizes()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, sizes)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles query execution error", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title, min_pages, max_pages FROM size").WillReturnError(errors.New("query error"))

		// Initialize repository
		repo := NewSizeRepository(db)

		// Call GetSizes
		sizes, err := repo.GetSizes()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, sizes)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
