package books

import (
	"net/http"

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

func (h *Handler) GetBookByCategory(c *gin.Context) {

	category := c.Query("category")
	// if (category == "") {
	// 	c.JSON(http.StatusBadRequest, g.H{"error": "category value is empty"})
	// }
	books, err := h.s.GetBookByCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
