package repositories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookRepository_GetBooks(t *testing.T) {
	t.Run("successfully fetches books", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewBookRepository(db)

		// Define query result
		rows := sqlmock.NewRows([]string{
			"book.id", "book.title", "book.year_published", "book.rating", "book.pages",
			"genre.id", "genre.title", "author.id", "author.first_name", "author.last_name",
		}).AddRow(
			1, "Book One", 2020, 4.5, 300,
			1, "Fiction", 1, "John", "Doe",
		).AddRow(
			2, "Book Two", 2021, 4.0, 250,
			2, "Non-Fiction", 2, "Jane", "Smith",
		)

		// Mock query
		mock.ExpectQuery("SELECT .* FROM book").
			WithArgs(300, 2020, 2021).
			WillReturnRows(rows)

		filters := map[string]interface{}{
			"min_pages": 300,
			"min_year":  2020,
			"max_year":  2021,
		}

		// Act
		books, err := repo.GetBooks(filters)

		// Assert
		assert.NoError(t, err)
		assert.Len(t, books, 2)

		assert.Equal(t, "Book One", books[0].Title)
		assert.Equal(t, "John", books[0].Author.FirstName)
		assert.Equal(t, "Fiction", books[0].Genre.Title)

		assert.Equal(t, "Book Two", books[1].Title)
		assert.Equal(t, "Jane", books[1].Author.FirstName)
		assert.Equal(t, "Non-Fiction", books[1].Genre.Title)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns empty list when no books match", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewBookRepository(db)

		// Mock query
		mock.ExpectQuery("SELECT .* FROM book").
			WithArgs(500, 2010, 2015).
			WillReturnRows(sqlmock.NewRows([]string{})) // No rows

		filters := map[string]interface{}{
			"min_pages": 500,
			"min_year":  2010,
			"max_year":  2015,
		}

		// Act
		books, err := repo.GetBooks(filters)

		// Assert
		assert.NoError(t, err)
		assert.Len(t, books, 0)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles query execution errors gracefully", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewBookRepository(db)

		// Mock query error
		mock.ExpectQuery("SELECT .* FROM book").
			WillReturnError(fmt.Errorf("query error"))

		filters := map[string]interface{}{
			"min_pages": 100,
		}

		// Act
		books, err := repo.GetBooks(filters)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, books)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
