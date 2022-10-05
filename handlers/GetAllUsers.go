package handlers

import (
	"go-gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAllUsers(ctx *gin.Context) {
	var users []models.User

	result := h.DB.Find(&users)

	if(result.Error != nil){
		log.Fatalln(result.Error)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
