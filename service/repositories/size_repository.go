package repositories

import (
	"database/sql"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
)

type SizeRepository struct {
	DB *sql.DB
}

// NewSizeRepository initializes the SizeRepository
func NewSizeRepository(db *sql.DB) *SizeRepository {
	return &SizeRepository{DB: db}
}

// GetSizes fetches all sizes from the database
func (r *SizeRepository) GetSizes() ([]models.Size, error) {
	query := "SELECT id, title, min_pages, max_pages FROM size"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sizes []models.Size
	for rows.Next() {
		var size models.Size
		var minPages, maxPages sql.NullInt32
		if err := rows.Scan(&size.ID, &size.Title, &minPages, &maxPages); err != nil {
			return nil, err
		}
		if minPages.Valid {
			min := int(minPages.Int32)
			size.MinPages = &min
		}
		if maxPages.Valid {
			max := int(maxPages.Int32)
			size.MaxPages = &max
		}
		sizes = append(sizes, size)
	}

	if len(sizes) == 0 || sizes == nil {
		sizes = []models.Size{}
	}

	return sizes, nil
}
