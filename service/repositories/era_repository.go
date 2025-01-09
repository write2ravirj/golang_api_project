package repositories

import (
	"database/sql"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
)

type ErasRepository struct {
	DB *sql.DB
}

// NewErasRepository initializes the ErasRepository
func NewErasRepository(db *sql.DB) *ErasRepository {
	return &ErasRepository{DB: db}
}

// GetEras fetches all eras from the database
func (r *ErasRepository) GetEras() ([]models.Era, error) {
	query := "SELECT id, title, min_year, max_year FROM era"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eras []models.Era
	for rows.Next() {
		var era models.Era
		var minYear, maxYear sql.NullInt32
		if err := rows.Scan(&era.ID, &era.Title, &minYear, &maxYear); err != nil {
			return nil, err
		}
		if minYear.Valid {
			min := int(minYear.Int32)
			era.MinYear = &min
		}
		if maxYear.Valid {
			max := int(maxYear.Int32)
			era.MaxYear = &max
		}
		eras = append(eras, era)
	}

	if len(eras) == 0 || eras == nil {
		eras = []models.Era{}
	}

	return eras, nil
}
