package main

import (
	"api/controllers"
	"api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/books", controllers.Findbooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
