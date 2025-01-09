package controllers

import (
	"encoding/json"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/Cors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	"net/http"
)

// GenreController handles HTTP requests for genres
type GenreController struct {
	Service services.GenreService
}

// NewGenreController initializes a new GenreController
func NewGenreController(service services.GenreService) *GenreController {
	return &GenreController{Service: service}
}

// GetGenres handles the HTTP GET request for fetching genres
func (c *GenreController) GetGenres(w http.ResponseWriter, r *http.Request) {
	Cors.EnableCORS(w)
	// Fetch genres from the service layer
	genres, err := c.Service.GetGenres()
	if err != nil {
		http.Error(w, "Error fetching genres", http.StatusInternalServerError)
		return
	}

	// Set content type as JSON and encode the genres list as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(genres); err != nil {
		http.Error(w, "Error encoding genres data", http.StatusInternalServerError)
	}
}
