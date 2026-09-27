package controllers

import (
	"net/http"
	"toko-buku-api1/config"
	"toko-buku-api1/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Books
	config.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBooks(c *gin.Context) {
	var input models.Books

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan", "data": input})
}
