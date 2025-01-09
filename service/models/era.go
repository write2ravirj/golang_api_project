package models

// Era represents an era in the system
type Era struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	MinYear *int   `json:"minYear,omitempty"`
	MaxYear *int   `json:"maxYear,omitempty"`
}
