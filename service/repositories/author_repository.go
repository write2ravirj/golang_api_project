package repositories

import (
	"database/sql"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
)

// DBQuerier defines an interface for the methods needed by AuthorRepository.
type DBQuerier interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// AuthorRepository is responsible for interacting with the authors table.
type AuthorRepository struct {
	DB DBQuerier
}

// NewAuthorRepository initializes a new AuthorRepository with the provided DBQuerier.
func NewAuthorRepository(db DBQuerier) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

// GetAuthors retrieves the list of authors from the database.
func (r *AuthorRepository) GetAuthors() ([]models.Author, error) {
	query := "SELECT id, first_name, last_name FROM author"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if len(authors) == 0 || authors == nil {
		authors = []models.Author{}
	}
	return authors, nil
}
