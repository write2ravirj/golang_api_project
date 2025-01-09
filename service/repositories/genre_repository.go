package repositories

import (
	"database/sql"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
)

type GenreRepository struct {
	DB *sql.DB
}

// NewGenreRepository initializes the GenreRepository
func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{DB: db}
}

// GetGenres fetches all genres from the database
func (r *GenreRepository) GetGenres() ([]models.Genre, error) {
	query := "SELECT id, title FROM genre"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Title); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	if len(genres) == 0 || genres == nil {
		genres = []models.Genre{}
	}

	return genres, nil
}
