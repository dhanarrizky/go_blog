package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dhanarrizky/go-blog/helper"
	"github.com/dhanarrizky/go-blog/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func hashPassword(password string) string {
	hashPsd, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panic(err.Error())
	}

	return string(hashPsd)
}

func verifyPassword(userPassword, providerPassword string) (string, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(providerPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}

	return msg, check
}

func UsersSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		var existUserUserName models.User
		if err := DB.Where("username = ?", user.UserName).First(&existUserUserName).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
			return
		}

		var existUserEmail models.User
		if err := DB.Where("email = ?", user.Email).First(&existUserEmail).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
			return
		}

		var existUserPhone models.User
		if err := DB.Where("phone = ?", user.Phone).First(&existUserPhone).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone number already exists"})
			return
		}

		strId := strconv.FormatUint(uint64(*&user.ID), 10)
		token, refreshToken, err := helper.GenerateJwtToken(*&user.UserName, *&user.Email, *&user.FirstName, *&user.LastName, strId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.Token = &token
		user.RefreshToken = &refreshToken

		db := DB.Create(&user)
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		if db.RowsAffected > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "users has been created successfully"})
		}
	}
}

func UsersLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var user models.User
		// var findUser models.User

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existUserUserName models.User
		if err := DB.Where("username = ?", user.UserName).First(&existUserUserName).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found"})
			return
		}

		msg, isValidPassword := verifyPassword(user.Password, existUserUserName.Password)

		if isValidPassword != true {
			c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		}

		strId := strconv.FormatUint(uint64(*&existUserUserName.ID), 10)
		token, refreshToken, err := helper.GenerateJwtToken(*&existUserUserName.UserName, *&existUserUserName.Email, *&existUserUserName.FirstName, *&existUserUserName.LastName, strId)
		helper.UpdateJwtToken(token, refreshToken, strId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existUserUserName)
	}
}
