package repositories

import (
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenreRepository_GetGenres(t *testing.T) {
	t.Run("successfully fetches genres", func(t *testing.T) {
		// Mock database and data
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "Science Fiction").
			AddRow(2, "Fantasy").
			AddRow(3, "Mystery")

		mock.ExpectQuery("SELECT id, title FROM genre").WillReturnRows(rows)

		// Initialize repository
		repo := NewGenreRepository(db)

		// Call GetGenres
		genres, err := repo.GetGenres()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, genres, 3)

		expected := []models.Genre{
			{ID: 1, Title: "Science Fiction"},
			{ID: 2, Title: "Fantasy"},
			{ID: 3, Title: "Mystery"},
		}
		assert.Equal(t, expected, genres)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles empty result set", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title FROM genre").WillReturnRows(sqlmock.NewRows(nil))

		// Initialize repository
		repo := NewGenreRepository(db)

		// Call GetGenres
		genres, err := repo.GetGenres()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, genres)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles query execution error", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, title FROM genre").WillReturnError(errors.New("query error"))

		// Initialize repository
		repo := NewGenreRepository(db)

		// Call GetGenres
		genres, err := repo.GetGenres()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, genres)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
