package handlers

import (
	"log"
	"net/http"

	"go-gin/models"

	"github.com/gin-gonic/gin"

	"go-gin/utils/validators"
)

func AddUser(ctx *gin.Context) {

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

	result, err := UserRepo.Save(&user)

	if err != nil {
		log.Fatalln(err.Error())
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": "Create Error Failed",
		})
		return
	}

	// Send a 201 created response
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User item created successfully!",
		"result":  result,
	})
}
