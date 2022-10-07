package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin/handlers"
)

func main() {
	r := gin.Default()

	//static file
	r.Static("/assets", "./assets")

	//GET method
	r.GET("/ping", getPing)

	//POST method
	r.POST("/post", postPing)

	//POST method with formdata
	r.POST("/formdata", adminMiddleware, getUserWithForm)

	//API group
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/heath", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "Service up!")
			})
		}
	}

	//upload single file
	r.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		ctx.SaveUploadedFile(file, "./assets/upload/"+file.Filename)
		ctx.String(http.StatusOK, file.Filename+" uploaded!")
	})

	//upload multiple files
	r.POST("/uploads", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			ctx.SaveUploadedFile(file, "./assets/upload/"+file.Filename)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "files uploaded!",
			"length":  len(files),
		})
	})

	//books service
	r.POST("/books", handlers.AddBook)
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.AddUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	//Run server
	r.Run(":3333")
}

func getPing(ctx *gin.Context) {

	ctx.String(http.StatusOK, "pong")
}

func postPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is a post request",
	})
}

func getUserWithForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.PostForm("age")
	ctx.String(http.StatusOK, "Name: "+name+" Age: "+age)
}

// Middleware
func adminMiddleware(ctx *gin.Context) {
	name := ctx.PostForm("name")

	if name != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not an admin",
		})
		ctx.Abort()
		return
	}

	ctx.Next()

}
