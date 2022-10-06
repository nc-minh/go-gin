package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"go-gin/models"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	var books []models.Book

	// if result := h.DB.Find(&books); result.Error != nil {
	// 	log.Fatalln(result.Error)
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func AddBook(ctx *gin.Context) {
	// Read to request body
	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	// if result := h.DB.Create(&book); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Send a 201 created response
	ctx.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Book item created successfully!",
		"resourceId": book.ID,
	})
}
