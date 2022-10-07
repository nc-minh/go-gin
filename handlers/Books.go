package handlers

import (
	"net/http"
	"strconv"

	"go-gin/models"

	"github.com/gin-gonic/gin"
)

//CRUD

func AddBook(ctx *gin.Context) {

	var book models.Book

	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	result, err := BookRepo.Save(&book)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Send a 201 created response
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Book item created successfully!",
		"result":  result,
	})
}

func GetBooks(ctx *gin.Context) {

	books, err := BookRepo.FindAll()

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Books retrieved successfully!",
		"data":    books,
	})
}

func GetBook(ctx *gin.Context) {

	var book models.Book

	id := ctx.Param("id")

	bookId, _ := strconv.Atoi(id)

	result, err := BookRepo.FindById(bookId)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	book = *result

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book retrieved successfully!",
		"data":    book,
	})
}

func UpdateBook(ctx *gin.Context) {

	var book models.Book

	id := ctx.Param("id")

	bookId, _ := strconv.Atoi(id)

	result, err := BookRepo.FindById(bookId)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	book = *result

	err = ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err = BookRepo.Update(&book)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book updated successfully!",
		"data":    result,
	})
}

func DeleteBook(ctx *gin.Context) {

	id := ctx.Param("id")

	bookId, _ := strconv.Atoi(id)

	book, err := BookRepo.FindById(bookId)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	err = BookRepo.Delete(book)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
