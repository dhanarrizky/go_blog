package controllers

import (
	"fmt"

	"github.com/dhanarrizky/go-blog/models"
)

func CreatePostControllers() {
	var post models.Post

	err := DB.Create(&post)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("Create post has been successfully")
	}
}

func ShowAllPostControllers() {
	var post []models.Post
	err := DB.Preload("Category").Find(&post)
	if err.Error != nil {
		panic(err.Error.Error())
	}

}

func ShowDetailePostControllers(id int) {
	var post models.Post
	err := DB.Preload("Category").Preload("User").Find(&post, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

}

func UpdatePostControllers(id int) {
	var post models.Post

	err := DB.Find(&post, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	err = DB.Save(&post)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("Create post has been successfully")
	}
}

func DeletePostControllers(id int) {
	var post models.Post
	err := DB.Find(&post, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	err = DB.Delete(&post)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("delete post has been successfully")
	}
}
