package repositories

import (
	"database/sql"
	"fmt"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/models"
	"strings"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) GetBooks(filters map[string]interface{}) ([]models.Book, error) {
	query := `
		SELECT book.id, book.title, book.year_published, book.rating, book.pages, 
		       genre.id, genre.title, author.id, author.first_name, author.last_name
		FROM book
		JOIN genre ON book.genre_id = genre.id 
		JOIN author ON book.author_id = author.id
		WHERE 1=1
	`

	params := []interface{}{}
	index := 1 // Placeholder index

	// Apply filters dynamically
	for key, value := range filters {
		switch key {
		case "author_id":
			// Ensure author_id is a string of comma-separated values or a slice of strings
			switch v := value.(type) {
			case []string:
				placeholders := []string{}
				ids := strings.Split(v[0], ",")
				for _, id := range ids {
					placeholders = append(placeholders, fmt.Sprintf("$%d", index))
					params = append(params, id)
					index++
				}
				query += fmt.Sprintf(" AND author.id IN (%s)", strings.Join(placeholders, ","))
			case string:
				// If it's a single string like "10,7", split it into individual IDs
				ids := strings.Split(v, ",")
				placeholders := []string{}
				for _, id := range ids {
					placeholders = append(placeholders, fmt.Sprintf("$%d", index))
					params = append(params, id)
					index++
				}
				query += fmt.Sprintf(" AND author.id IN (%s)", strings.Join(placeholders, ","))
			}
		case "genre_id":
			// Same handling as author_id for genre
			switch v := value.(type) {
			case []string:
				ids := strings.Split(v[0], ",")
				placeholders := []string{}
				for _, id := range ids {
					placeholders = append(placeholders, fmt.Sprintf("$%d", index))
					params = append(params, id)
					index++
				}
				query += fmt.Sprintf(" AND genre.id IN (%s)", strings.Join(placeholders, ","))
			case string:
				// If it's a single string like "10,7", split it into individual IDs
				ids := strings.Split(v, ",")
				placeholders := []string{}
				for _, id := range ids {
					placeholders = append(placeholders, fmt.Sprintf("$%d", index))
					params = append(params, id)
					index++
				}
				query += fmt.Sprintf(" AND genre.id IN (%s)", strings.Join(placeholders, ","))
			}
		case "min_pages":
			query += fmt.Sprintf(" AND book.pages >= $%d", index)
			params = append(params, value)
			index++
		case "max_pages":
			query += fmt.Sprintf(" AND book.pages <= $%d", index)
			params = append(params, value)
			index++
		case "min_year":
			query += fmt.Sprintf(" AND book.year_published >= $%d", index)
			params = append(params, value)
			index++
		case "max_year":
			query += fmt.Sprintf(" AND book.year_published <= $%d", index)
			params = append(params, value)
			index++
		}
	}

	query += " ORDER BY book.rating DESC"

	if limit, ok := filters["limit"]; ok {
		query += fmt.Sprintf(" LIMIT $%d", index)
		params = append(params, limit)
		index++
	}

	// Execute the query
	rows, err := r.DB.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Parse results
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.YearPublished, &book.Rating, &book.Pages,
			&book.Genre.ID, &book.Genre.Title,
			&book.Author.ID, &book.Author.FirstName, &book.Author.LastName)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		books = append(books, book)
	}
	if len(books) == 0 || books == nil {
		books = []models.Book{}
	}
	return books, nil
}
