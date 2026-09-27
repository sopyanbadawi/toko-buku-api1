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

func GetBookById(c *gin.Context) {
	var book models.Books
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditenmukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
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

func UpdateBook(c *gin.Context) {
	var book models.Books
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	var input models.Books
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diperbarui", "data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Books
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}
