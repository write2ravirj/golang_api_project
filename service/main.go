package main

import (
	"database/sql"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/controllers"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/dbConnect"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/repositories"
	"git.codesubmit.io/nesto/back-end-code-challenge-tsvihr/service/services"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	db = dbConnect.InitDB()
	defer db.Close()

	// Initialize repositories, services, and controllers
	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	// Initialize repositories, services, and controllers
	authorRepo := repositories.NewAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepo)
	authorController := controllers.NewAuthorController(authorService)

	// Initialize repositories, services, and controllers
	genreRepo := repositories.NewGenreRepository(db)
	genreService := services.NewGenreService(genreRepo)
	genreController := controllers.NewGenreController(genreService)

	// Initialize repositories, services, and controllers
	sizeRepo := repositories.NewSizeRepository(db)
	sizeService := services.NewSizeService(sizeRepo)
	sizeController := controllers.NewSizeController(sizeService)

	// Initialize repositories, services, and controllers
	eraRepo := repositories.NewErasRepository(db)
	eraService := services.NewErasService(eraRepo)
	eraController := controllers.NewErasController(eraService)

	// Register routes
	http.HandleFunc("/api/v1/books", bookController.GetBooks)
	http.HandleFunc("/api/v1/authors", authorController.GetAuthors)
	http.HandleFunc("/api/v1/genres", genreController.GetGenres)
	http.HandleFunc("/api/v1/sizes", sizeController.GetSizes)
	http.HandleFunc("/api/v1/eras", eraController.GetEras)

	log.Println("Server is running on http://localhost:5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
