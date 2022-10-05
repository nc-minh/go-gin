package handlers

import (
	"fmt"
	"log"
	"net/http"

	"go-gin/models"

	"github.com/gin-gonic/gin"

	"go-gin/utils/validators"
)

func (h handler) AddUser(ctx *gin.Context) {

	var user models.User

	err := ctx.ShouldBindJSON(&user)

	var validationError error = validators.Validate(&user)

	if validationError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validationError.Error(),
		})
		return
	}

	if err != nil {
		log.Fatalln(err.Error())
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": "Create Error Failed",
		})
		return
	}

	// Append to the Books table
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	ctx.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "User item created successfully!",
		"resourceId": user.ID,
	})
}
