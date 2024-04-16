package books

type Book struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Pages         int64   `json:"pages"`
	PublishedDate string  `json:"publishedDate"`
}

type Repository interface {
	CreateBook(book *Book) (*Book, error)
	GetAllBooks() ([]Book, error)
	GetBooks(category string, title string) ([]Book, error)
	GetBookByID(id int64) (Book, error)
	UpdateBook(id int64, book Book) (Book, error)
	DeleteBook(id int64) (int64, error)
}

type CreateBookRequest struct {
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Pages         int64   `json:"pages"`
	PublishedDate string  `json:"publishedDate"`
}
