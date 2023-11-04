package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/dhanarrizky/go-blog/database"
	"github.com/dhanarrizky/go-blog/models"
)

var DB = database.ConDB()

func CreateUsersControllers() {
	var users models.User

	users.UserName = "dhanar"
	users.Email = "dhanar@gmail.com"
	users.Password = "dhanar"
	users.FirstName = "dhanar"
	users.LastName = "rizky"
	users.Phone = 84342344325555
	users.Gender = true
	users.Address = "wonogiri"
	dob := time.Date(2001, time.Month(8), 23, 0, 0, 0, 0, time.UTC)
	users.BirthDate = dob
	users.BirthPlace = "wonogiri"

	result := DB.Create(&users)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	if result.RowsAffected > 0 {
		log.Println("Created data has been successfully")
	}
}

func ShowAllUserControllers() {
	var users []models.User
	err := DB.Find(&users)

	if err.Error != nil {
		panic(err.Error.Error())
	}

	for i, val := range users {
		fmt.Println(i, " => ")
		fmt.Println("\t", val.ID)
		fmt.Println("\t", val.FirstName)
		fmt.Println("\t", val.LastName)
	}
}

func ShowDetaileControllers(id int) {
	var users models.User
	err := DB.Find(&users, id)

	if err.Error != nil {
		fmt.Println("user not found")
	}
	fmt.Println("UserName : ", users.UserName)
	fmt.Println("Email : ", users.Email)
	fmt.Println("Password : ", users.Password)
	fmt.Println("FirstName : ", users.FirstName)
	fmt.Println("LastName : ", users.LastName)
	fmt.Println("Phone : ", users.Phone)
	fmt.Println("Gender : ", users.Gender)
	fmt.Println("Address : ", users.Address)
	fmt.Println("BirthDate : ", users.BirthDate)
	fmt.Println("BirthPlace : ", users.BirthPlace)
}

func UpdateUserControllers(id int) {
	var users models.User
	err := DB.First(&users, id)
	if err.Error != nil {
		fmt.Println("user not found")
	}

	fmt.Println("\t", users.ID)
	users.Email = "yoi@gmail.com"
	users.Password = "nonsafsafsae"
	users.FirstName = "fsfsaf"
	users.LastName = "rizsfsafsafsafky"
	users.Gender = true
	users.Address = "fsfsafsafasfsfdfdfs"
	dob := time.Date(2003, time.Month(2), 3, 0, 0, 0, 0, time.UTC)
	users.BirthDate = dob
	users.BirthPlace = "wonogiri"

	fmt.Println("\t", users.ID)
	fmt.Println("\t", users.FirstName)
	fmt.Println("\t", users.LastName)
	fmt.Println("\t", users.BirthPlace)
	fmt.Println("\t", users.BirthDate)
	fmt.Println("\t", users.Address)

	DB.Save(&users)

	if err.RowsAffected > 0 {
		fmt.Println("updated user has been successfully")
	}
}

func DeleteUserControllers(id int) {
	var users models.User
	err := DB.Delete(&users, id)
	if err.Error != nil {
		fmt.Println("user not found")
	}

	if err.RowsAffected > 0 {
		fmt.Println("deleted user has been successfully")
	}
}
