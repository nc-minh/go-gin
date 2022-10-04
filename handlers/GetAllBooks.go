package handlers

import (
	"fmt"
	models "go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAllBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
