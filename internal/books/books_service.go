package books

type service struct {
	Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) CreateBook(req *CreateBookRequest) (*Book, error) {

	book := &Book{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Price:         req.Price,
		Pages:         req.Pages,
		PublishedDate: req.PublishedDate,
	}

	createdBook, err := s.Repository.CreateBook(book)

	if err != nil {
		return nil, err
	}

	return createdBook, nil
}

func (s *service) GetAllBooks() ([]Book, error) {
	allBooks, err := s.Repository.GetAllBooks()

	if err != nil {
		return nil, err
	}

	return allBooks, nil
}

func (s *service) GetBooks(category string, title string) ([]Book, error) {
	allBooks, err := s.Repository.GetBooks(category, title)

	if err != nil {
		return nil, err
	}

	if allBooks == nil {
		return nil, nil
	}

	return allBooks, nil
}

func (s *service) GetBookByID(id int64) (Book, error) {
	book, err := s.Repository.GetBookByID(id)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (s *service) UpdateBook(id int64, book Book) (Book, error) {
	book, err := s.Repository.UpdateBook(id, book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (s *service) DeleteBook(id int64) (int64, error) {

	rows, err := s.Repository.DeleteBook(id)
	if err != nil {
		return 0, err
	}

	return rows, nil
}
