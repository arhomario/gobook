package v1

import (
	"gobook/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type StoreBookInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateBookInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetBooks(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func StoreBook(ctx *gin.Context) {

	var input StoreBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	book := models.Book{Title: input.Title, Content: input.Content}

	db := ctx.MustGet("db").(*gorm.DB)
	db.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func EditBook(ctx *gin.Context) {
	var book models.Book

	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func UpdateBook(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not Found",
		})
		return
	}

	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedInput models.Book
	updatedInput.Title = input.Title
	updatedInput.Content = input.Content

	db.Model(&book).Update(updatedInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func DeleteBook(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not Found",
		})
		return
	}

	db.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": true,
	})

}
