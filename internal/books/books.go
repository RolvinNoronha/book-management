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
	GetBookByCategory(category string) ([]Book, error)
}

type CreateBookRequest struct {
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Pages         int64   `json:"pages"`
	PublishedDate string  `json:"publishedDate"`
}
