package books

import (
	"database/sql"
)

type DBTX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	// PrepareContext(context.Context, string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBook(book *Book) (*Book, error) {

	query := "INSERT INTO books(title, author, category, price, pages, publishedDate) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, book.Title, book.Author, book.Category, book.Price, book.Pages, book.PublishedDate)

	if err != nil {
		return nil, err
	}

	bookId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.ID = int64(bookId)
	return book, nil
}

func (r *repository) GetAllBooks() ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var allBooks []Book
	defer rows.Close()

	for rows.Next() {
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Price, &book.Pages, &book.PublishedDate)

		allBooks = append(allBooks, book)
	}

	return allBooks, nil
}

func (r *repository) GetBookByCategory(category string) ([]Book, error) {
	query := "SELECT * FROM books WHERE category = ?"
	rows, err := r.db.Query(query, category)

	if err != nil {
		return nil, err
	}

	var allBooks []Book
	defer rows.Close()

	for rows.Next() {
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Price, &book.Pages, &book.PublishedDate)

		allBooks = append(allBooks, book)
	}

	if len(allBooks) == 0 {
		return nil, nil
	}

	return allBooks, nil
}
