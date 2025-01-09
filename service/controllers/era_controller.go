package controllers

import (
	"encoding/json"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/Cors"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	"net/http"
)

// ErasController handles HTTP requests for eras
type ErasController struct {
	Service services.ErasService
}

// NewErasController initializes a new ErasController
func NewErasController(service services.ErasService) *ErasController {
	return &ErasController{Service: service}
}

// GetEras handles the HTTP GET request for fetching eras
func (c *ErasController) GetEras(w http.ResponseWriter, r *http.Request) {
	Cors.EnableCORS(w)
	// Fetch eras from the service layer
	eras, err := c.Service.GetEras()
	if err != nil {
		http.Error(w, "Error fetching eras", http.StatusInternalServerError)
		return
	}

	// Set content type as JSON and encode the eras list as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(eras); err != nil {
		http.Error(w, "Error encoding eras data", http.StatusInternalServerError)
	}
}
