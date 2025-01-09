package models

type Size struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	MinPages *int   `json:"minPages,omitempty"`
	MaxPages *int   `json:"maxPages,omitempty"`
}
