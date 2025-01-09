package repositories

import (
	"errors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorRepository_GetAuthors(t *testing.T) {
	t.Run("successfully fetches authors", func(t *testing.T) {
		// Mock database and data
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
			AddRow(1, "John", "Doe").
			AddRow(2, "Jane", "Smith")

		mock.ExpectQuery("SELECT id, first_name, last_name FROM author").WillReturnRows(rows)

		// Initialize repository
		repo := NewAuthorRepository(db)

		// Call GetAuthors
		authors, err := repo.GetAuthors()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, authors, 2)

		expected := []models.Author{
			{ID: 1, FirstName: "John", LastName: "Doe"},
			{ID: 2, FirstName: "Jane", LastName: "Smith"},
		}
		assert.Equal(t, expected, authors)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles empty result set", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, first_name, last_name FROM author").WillReturnRows(sqlmock.NewRows(nil))

		// Initialize repository
		repo := NewAuthorRepository(db)

		// Call GetAuthors
		authors, err := repo.GetAuthors()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, authors)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles query execution error", func(t *testing.T) {
		// Mock database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT id, first_name, last_name FROM author").WillReturnError(errors.New("query error"))

		// Initialize repository
		repo := NewAuthorRepository(db)

		// Call GetAuthors
		authors, err := repo.GetAuthors()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, authors)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
