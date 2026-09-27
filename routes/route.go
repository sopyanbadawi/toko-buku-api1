package routes

import (
	"toko-buku-api1/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/books", controllers.GetBooks)
		api.GET("/books/:id", controllers.GetBookById)
		api.POST("/books", controllers.CreateBooks)
		api.PUT("/books/:id", controllers.UpdateBook)
		api.DELETE("/books/:id", controllers.DeleteBook)

	}

	return r
}
