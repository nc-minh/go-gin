package handlers

import (
	"net/http"
	"strconv"

	"go-gin/models"

	"github.com/gin-gonic/gin"

	"go-gin/utils/validators"
)

func GetAllUsers(ctx *gin.Context) {

	result, err := UserRepo.FindAll()

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, _ := strconv.Atoi(id)

	result, err := UserRepo.FindById(userId)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, _ := strconv.Atoi(id)

	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	user.ID = uint(userId)

	_, err = UserRepo.Update(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{

		"message": "User updated successfully",

		"data": user,
	})

}

func DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	userId, _ := strconv.Atoi(id)

	user, err := UserRepo.FindById(userId)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	err = UserRepo.Delete(user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Error: ": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

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
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := UserRepo.Save(&user)

	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
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
