package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {

	result, err := UserRepo.FindAll()

	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
