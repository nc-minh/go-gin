package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "go-gin/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddBook(ctx *gin.Context) {
	// Read to request body
	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	ctx.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Book item created successfully!",
		"resourceId": book.ID,
	})
}
