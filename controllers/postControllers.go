package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhanarrizky/go-blog/models"
	"github.com/gin-gonic/gin"
)

func CreatePostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var post models.Post

		if err := c.Bind(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := DB.Create(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()

		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, post)
			// fmt.Println("Create post has been successfully")
		}
	}
}

func ShowAllPostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var post models.Post

		posts := map[string]interface{}{}
		db := DB.Model(&post).Preload("Category").First(&posts)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		count := len(posts)
		groupJson := gin.H{
			"count": count,
			"post":  posts,
		}

		c.JSON(http.StatusOK, groupJson)
	}
}

func ShowDetailePostControllers(id int) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var post models.Post

		postId := c.Param("id")
		db := DB.Preload("Category").Preload("User").Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, post)
	}
}

func UpdatePostControllers(id int) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var post models.Post

		postId := c.Param("id")
		db := DB.Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		db = DB.Save(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		if db.RowsAffected > 0 {
			// fmt.Println("Create post has been successfully")
			c.JSON(http.StatusOK, post)
		}
	}
}

func DeletePostControllers(id int) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var post models.Post

		postId := c.Param("id")
		db := DB.Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		db = DB.Delete(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		defer cancel()
		if db.RowsAffected > 0 {
			// fmt.Println("delete post has been successfully")
			c.JSON(http.StatusOK, gin.H{"message": "delete post has been successfully"})
		}
	}
}
