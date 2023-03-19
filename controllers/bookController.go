package controllers

import (
	"api/db"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find all books
func Findbooks(c *gin.Context) {
	var books []models.Book
	db.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Create a book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Find a book
func FindBook(c *gin.Context) {
	var book models.Book
	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update a book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{Title: input.Title, Author: input.Author}

	db.DB.Model(&book).Updates(&updateBook)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := db.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
