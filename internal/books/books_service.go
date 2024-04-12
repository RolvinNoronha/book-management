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

func (s *service) GetBookByCategory(category string) ([]Book, error) {
	allBooks, err := s.Repository.GetBookByCategory(category)

	if err != nil {
		return nil, err
	}

	if allBooks == nil {
		return nil, nil
	}

	return allBooks, nil
}
