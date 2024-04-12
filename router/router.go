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
	r.GET("/get-books", bookHandler.GetBookByCategory)
}

func Start(addr string) error {
	return r.Run(addr)
}
