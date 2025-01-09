package models

type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	YearPublished int     `json:"yearPublished"`
	Rating        float64 `json:"rating"`
	Pages         int     `json:"pages"`
	Genre         Genre   `json:"genre"`
	Author        Author  `json:"author"`
}
