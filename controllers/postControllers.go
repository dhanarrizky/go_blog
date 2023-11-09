package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dhanarrizky/go-blog/helper"
	"github.com/dhanarrizky/go-blog/models"
	"github.com/gin-gonic/gin"
)

func CreatePostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		var categories models.Categories
		var users models.User

		userId := c.GetString("uId")
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		if err := c.Bind(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dbUser := DB.Find(&users, userId)
		if dbUser.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbUser.Error.Error()})
			return
		}

		dbCategory := DB.Find(&categories, post.CategoryID)
		if dbCategory.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbUser.Error.Error()})
			return
		}

		intUserId, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		post.UserID = uint(intUserId)
		post.User = users
		post.Category = categories
		db := DB.Create(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		users.Posts = append(users.Posts, post)
		dbUser.Save(users)
		if dbUser.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbUser.Error.Error()})
			return
		}

		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, post)
			// fmt.Println("Create post has been successfully")
		}
	}
}

func ShowAllPostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post []models.Post
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		db := DB.Preload("Category").Find(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		count := len(post)
		groupJson := gin.H{
			"count": count,
			"post":  post,
		}

		c.JSON(http.StatusOK, groupJson)
	}
}

func ShowDetailePostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		postId := c.Param("id")

		db := DB.Preload("Category").Preload("User").Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

func UpdatePostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post

		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		postId := c.Param("id")
		db := DB.Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		stringUserId := strconv.Itoa(int(post.UserID))
		if err := helper.AdminValidate(c, stringUserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db = DB.Save(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		if db.RowsAffected > 0 {
			// fmt.Println("Create post has been successfully")
			c.JSON(http.StatusOK, post)
		}
	}
}

func DeletePostControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		postId := c.Param("id")
		db := DB.Find(&post, postId)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		stringUserId := strconv.Itoa(int(post.UserID))
		if err := helper.AdminValidate(c, stringUserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db = DB.Delete(&post)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		if db.RowsAffected > 0 {
			// fmt.Println("delete post has been successfully")
			c.JSON(http.StatusOK, gin.H{"message": "delete post has been successfully"})
		}
	}
}
