package repositories

import (
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErasRepository_GetEras(t *testing.T) {
	t.Run("successfully fetches eras", func(t *testing.T) {
		// Mock database and data
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "title", "min_year", "max_year"}).
			AddRow(1, "Renaissance", 1300, 1600).
			AddRow(2, "Baroque", 1600, 1750).
			AddRow(3, "Modern Era", nil, 2000)

		mock.ExpectQuery("SELECT id, title, min_year, max_year FROM era").WillReturnRows(rows)

		// Initialize repository
		repo := NewErasRepository(db)

		// Call GetEras
		eras, err := repo.GetEras()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, eras, 3)

		expected := []models.Era{
			{ID: 1, Title: "Renaissance", MinYear: intPtr(1300), MaxYear: intPtr(1600)},
			{ID: 2, Title: "Baroque", MinYear: intPtr(1600), MaxYear: intPtr(1750)},
			{ID: 3, Title: "Modern Era", MinYear: nil, MaxYear: intPtr(2000)},
		}
		assert.Equal(t, expected, eras)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles empty result set", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title, min_year, max_year FROM era").WillReturnRows(sqlmock.NewRows(nil))

		// Initialize repository
		repo := NewErasRepository(db)

		// Call GetEras
		eras, err := repo.GetEras()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, eras)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles query execution error", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title, min_year, max_year FROM era").WillReturnError(errors.New("query error"))

		// Initialize repository
		repo := NewErasRepository(db)

		// Call GetEras
		eras, err := repo.GetEras()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, eras)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

// intPtr is a helper function to return a pointer to an int.
func intPtr(i int) *int {
	return &i
}
