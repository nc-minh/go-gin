package handlers

import (
	"go-gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAllBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		log.Fatalln(result.Error)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
