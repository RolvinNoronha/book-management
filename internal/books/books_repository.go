package books

import (
	"database/sql"
	"fmt"
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

func (r *repository) GetBooks(category string, title string) ([]Book, error) {
	var rows *sql.Rows
	var err error
	if category != "" && title != "" {
		query := "SELECT * FROM books WHERE category = ? AND title LIKE ?"
		title = "%" + title + "%"
		rows, err = r.db.Query(query, category, title)

	} else if title != "" {
		query := "SELECT * FROM books WHERE title LIKE ?"
		title = "%" + title + "%"
		rows, err = r.db.Query(query, title)
	} else if category != "" {
		query := "SELECT * FROM books WHERE category = ?"
		rows, err = r.db.Query(query, category)
	}

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

func (r *repository) DeleteBook(id int64) (int64, error) {
	query := "DELETE FROM books WHERE id = ?"
	result, err := r.db.Exec(query, id)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (r *repository) GetBookByID(id int64) (Book, error) {
	query := "SELECT * FROM books WHERE id = ?"
	rows, err := r.db.Query(query, id)

	if err != nil {
		return Book{}, err
	}

	defer rows.Close()
	var book Book
	for rows.Next() {
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Price, &book.Pages, &book.PublishedDate)
	}

	return book, nil
}

func (r *repository) UpdateBook(id int64, book Book) (Book, error) {
	query := "UPDATE books SET "

	var args []interface{}
	if book.Title != "" {
		query += "title = ?, "
		args = append(args, book.Title)
	}

	if book.Author != "" {
		query += "author = ?, "
		args = append(args, book.Author)
	}

	if book.Category != "" {
		query += "category = ?, "
		args = append(args, book.Category)
	}

	if book.Price != 0 {
		query += "price = ?, "
		args = append(args, book.Price)
	}

	if book.Pages != 0 {
		query += "pages = ?, "
		args = append(args, book.Pages)
	}

	if book.PublishedDate != "" {
		query += "publishedDate = ?, "
		args = append(args, book.PublishedDate)
	}

	// Remove trailing comma and space
	query = query[:len(query)-2]

	query += " WHERE id = ?"
	args = append(args, id)

	fmt.Print(query)
	fmt.Print(args...)

	_, err := r.db.Exec(query, args...)

	if err != nil {
		return Book{}, err
	}

	updatedBook, err := r.GetBookByID(id)

	if err != nil {
		return Book{}, nil
	}

	return updatedBook, nil
}
