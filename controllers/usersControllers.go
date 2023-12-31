package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dhanarrizky/go-blog/database"
	"github.com/dhanarrizky/go-blog/helper"
	"github.com/dhanarrizky/go-blog/models"
	"github.com/gin-gonic/gin"
)

var DB = database.ConDB()

func ShowAllUserControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		if err := helper.AdminValidate(c, ""); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		err := DB.Find(&users)
		if err.Error != nil {
			log.Println(err.Error.Error())
			return
		}

		count := len(users)
		groupJson := gin.H{
			"count": count,
			"users": *&users,
		}

		c.JSON(http.StatusOK, groupJson)
	}
}

func ShowUserDetaileControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		userId := c.Param("id")
		err := DB.Find(&user, userId)

		if err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "user not found", "error": err.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := helper.AdminValidate(c, c.GetString("id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		userId := c.Param("id")
		err := DB.Find(&user, userId)
		if err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "user not found", "error": err.Error.Error()})
			return
		}

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		DB.Save(&user)

		if err.RowsAffected > 0 {
			// fmt.Println("updated user has been successfully")
			c.JSON(http.StatusOK, user)
		}
	}
}

func DeleteUserControllers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users models.User
		if err := helper.AdminValidate(c, c.GetString("id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		userId := c.Param("id")
		err := DB.Delete(&users, userId)
		if err.Error != nil {
			// fmt.Println("user not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error.Error()})
			return
		}

		if err.RowsAffected > 0 {
			// fmt.Println("deleted user has been successfully")
			c.JSON(http.StatusOK, gin.H{"message": "deleted user has been successfully"})
		}
	}
}
