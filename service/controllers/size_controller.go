package controllers

import (
	"encoding/json"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/Cors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	"net/http"
)

// SizeController handles HTTP requests for sizes
type SizeController struct {
	Service services.SizeService
}

// NewSizeController initializes a new SizeController
func NewSizeController(service services.SizeService) *SizeController {
	return &SizeController{Service: service}
}

// GetSizes handles the HTTP GET request for fetching sizes
func (c *SizeController) GetSizes(w http.ResponseWriter, r *http.Request) {
	Cors.EnableCORS(w)
	// Fetch sizes from the service layer
	sizes, err := c.Service.GetSizes()
	if err != nil {
		http.Error(w, "Error fetching sizes", http.StatusInternalServerError)
		return
	}

	// Set content type as JSON and encode the sizes list as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sizes); err != nil {
		http.Error(w, "Error encoding sizes data", http.StatusInternalServerError)
	}
}
