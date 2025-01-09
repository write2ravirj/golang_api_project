package controllers

import (
	"encoding/json"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/Cors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	"net/http"
)

type AuthorController struct {
	Service services.AuthorService
}

// NewAuthorController initializes a new AuthorController
func NewAuthorController(service services.AuthorService) *AuthorController {
	return &AuthorController{Service: service}
}

// GetAuthors handles the HTTP GET request for fetching authors
func (c *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	Cors.EnableCORS(w)
	// Fetch authors from the service layer

	authors, err := c.Service.GetAuthors()
	if err != nil {
		http.Error(w, "Error fetching authors", http.StatusInternalServerError)
		return
	}

	// Set content type as JSON and encode the authors list as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		http.Error(w, "Error encoding authors data", http.StatusInternalServerError)
	}
}
