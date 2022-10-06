package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "go-gin/databases"
	"go-gin/handlers"

	repoimpl "go-gin/repository/impl"
)

var UserRepo = repoimpl.NewUserRepo(db.Init())

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

	//init db

	//books service
	r.POST("/books", handlers.AddBook)
	r.GET("/books", handlers.GetAllBooks)
	r.GET("/users", handlers.GetAllUsers)
	r.POST("/users", handlers.AddUser)

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

// func getUser(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	ctx.String(http.StatusOK, "User ID: "+id)
// }

// func getUserWithQueryString(ctx *gin.Context) {
// 	name := ctx.Query("name")
// 	age := ctx.Query("age")
// 	ctx.String(http.StatusOK, "Name: "+name+" Age: "+age)
// }

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
