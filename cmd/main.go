package main

import (
	"log"

	"github.com/RolvinNoronha/book-management/db"
	"github.com/RolvinNoronha/book-management/internal/books"
	"github.com/RolvinNoronha/book-management/router"
)

func main() {
	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize the database %s", err)
	}

	log.Print("successfully initialized and connected to the database")

	bookRepository := books.NewRepository(dbConn.GetDb())
	bookService := books.NewService(bookRepository)
	bookHandler := books.NewHandler(*bookService)

	router.InitHandler(bookHandler)

	router.Start("0.0.0.0:8080")
}
