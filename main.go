package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Clean Code", Author: "Uncle bob", Quantity: 2},
	{ID: "2", Title: "Clean Architecture", Author: "Uncle bob", Quantity: 5},
	{ID: "3", Title: "Pragmatic Programmer", Author: "Richard", Quantity: 12},
	{ID: "4", Title: "Refactoring", Author: "Stephen", Quantity: 22},
}

func getBookById(id string) (*Book, error) {
	for index, value := range books {
		if value.ID == id {
			return &books[index], nil
		}
	}
	return nil, errors.New("Book not found")

}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func getBook(context *gin.Context) {
	bookId := context.Param("id")
	book, err := getBookById(bookId)
	if err != nil {
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func createBook(context *gin.Context) {
	var newBook Book
	if err := context.BindJSON(&newBook); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.POST("/books/:id", getBook)
	router.Run("localhost:6350")
}
