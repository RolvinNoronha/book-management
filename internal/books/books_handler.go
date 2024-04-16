package books

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s service
}

func NewHandler(s service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) CreateBook(c *gin.Context) {
	var book CreateBookRequest

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := h.s.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &Book{
		ID:            res.ID,
		Title:         res.Title,
		Author:        res.Author,
		Category:      res.Category,
		Price:         res.Price,
		Pages:         res.Pages,
		PublishedDate: res.PublishedDate,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetAllBooks(c *gin.Context) {

	books, err := h.s.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBooks(c *gin.Context) {

	category := c.Query("category")
	title := c.Query("title")
	// if (category == "") {
	// 	c.JSON(http.StatusBadRequest, g.H{"error": "category value is empty"})
	// }
	books, err := h.s.GetBooks(category, title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no books with the title or category exist."})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBookByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.s.GetBookByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no book with the id exist."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := h.s.UpdateBook(id, book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedBook.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no book with id exist."})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	idValue := c.Param("id")

	if idValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id cannot be empty"})
		return
	}

	id, err := strconv.Atoi(idValue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = h.s.DeleteBook(int64(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
