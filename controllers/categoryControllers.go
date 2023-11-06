package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhanarrizky/go-blog/models"
	"github.com/gin-gonic/gin"
)

func CreateCategoryControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Categories
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		if err := c.Bind(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := DB.Create(&category)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()

		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, category)
		}
	}
}

func ShowAllCategoryControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Categories
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		categories := map[string]interface{}{}
		db := DB.Model(&category).First(&categories)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		count := len(categories)
		groupJson := gin.H{
			"count":      count,
			"categories": categories,
		}

		c.JSON(http.StatusOK, groupJson)

	}
}

func ShowDetaileCategoryControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Categories
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		ctgId := c.Param("id")
		db := DB.Find(&category, ctgId)

		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, category)
	}
}

func UpdateCategoryControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Categories
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		ctgId := c.Param("id")
		db := DB.Find(&category, ctgId)

		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		}

		if err := c.Bind(&category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer cancel()
		DB.Save(&category)

		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, category)
		}
	}
	// fmt.Println("updated category has been successfully")
}

func DeleteCategoryControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Categories
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		ctgId := c.Param("id")
		db := DB.Delete(&category, ctgId)

		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "deleted category has been successfully"})
		}

	}
	// fmt.Println("deleted category has been successfully")
}
