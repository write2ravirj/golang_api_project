package controllers

import (
	"encoding/json"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/Cors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type BookController struct {
	Service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{Service: service}
}

// Helper function to sanitize and validate integer inputs
func sanitizeIntInput(param string, minValue, maxValue int) (int, bool) {
	if param == "" {
		return 0, false
	}
	value, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}
	// Validate if within acceptable range
	if value < minValue || value > maxValue {
		return 0, false
	}
	return value, true
}

// Helper function to sanitize comma-separated string inputs
func sanitizeCSVInput(param string) ([]string, bool) {
	if param == "" {
		return nil, false
	}
	// Trim spaces and split by commas
	values := strings.FieldsFunc(param, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ','
	})
	if len(values) == 0 {
		return nil, false
	}
	return values, true
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	Cors.EnableCORS(w)
	// Parse query parameters into a filters map
	filters := make(map[string]interface{})

	// Sanitize and validate authors parameter
	if authors := r.URL.Query().Get("authors"); authors != "" {
		ids, valid := sanitizeCSVInput(authors)
		if valid {
			filters["author_id"] = ids
		}
	}

	// Sanitize and validate genres parameter
	if genres := r.URL.Query().Get("genres"); genres != "" {
		ids, valid := sanitizeCSVInput(genres)
		if valid {
			filters["genre_id"] = ids
		}
	}

	// Sanitize and validate min-pages parameter
	if minPages := r.URL.Query().Get("min-pages"); minPages != "" {
		if min, valid := sanitizeIntInput(minPages, 1, 10000); valid {
			filters["min_pages"] = min
		}
	}

	// Sanitize and validate max-pages parameter
	if maxPages := r.URL.Query().Get("max-pages"); maxPages != "" {
		if max, valid := sanitizeIntInput(maxPages, 1, 10000); valid {
			filters["max_pages"] = max
		}
	}

	// Sanitize and validate min-year parameter
	if minYear := r.URL.Query().Get("min-year"); minYear != "" {
		if min, valid := sanitizeIntInput(minYear, 1000, 2100); valid {
			filters["min_year"] = min
		}
	}

	// Sanitize and validate max-year parameter
	if maxYear := r.URL.Query().Get("max-year"); maxYear != "" {
		if max, valid := sanitizeIntInput(maxYear, 1000, 2100); valid {
			filters["max_year"] = max
		}
	}

	// Sanitize and validate limit parameter
	if limit := r.URL.Query().Get("limit"); limit != "" {
		if limitLen, valid := sanitizeIntInput(limit, 1, 100); valid {
			filters["limit"] = limitLen
		}
	}

	books, err := c.Service.GetBooks(filters)
	if err != nil {
		http.Error(w, "Error fetching books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Error encoding books data", http.StatusInternalServerError)
	}
}
