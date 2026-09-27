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
		api.POST("/books", controllers.CreateBooks)
	}

	return r
}
