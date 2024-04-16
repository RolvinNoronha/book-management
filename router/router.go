package router

import (
	"github.com/RolvinNoronha/book-management/internal/books"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitHandler(bookHandler *books.Handler) {

	r = gin.Default()

	r.POST("/create-book", bookHandler.CreateBook)
	r.GET("/get-all-books", bookHandler.GetAllBooks)
	r.GET("/get-books", bookHandler.GetBooks)
	r.DELETE("/delete-book/:id", bookHandler.DeleteBook)
	r.GET("/get-book/:id", bookHandler.GetBookByID)
	r.POST("/update-book/:id", bookHandler.UpdateBook)
}

func Start(addr string) error {
	return r.Run(addr)
}
