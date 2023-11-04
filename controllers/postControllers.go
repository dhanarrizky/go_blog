package controllers

import (
	"fmt"

	"github.com/dhanarrizky/go-blog/models"
)

func CreatePostControllers() {
	var post models.Post

	post.UserID = 1
	post.Title = "hallo"
	post.CategoryID = 1
	post.Img = "hallo iugbdfsoiusagfoihsaofhai"
	post.Description = "hallo iugbdfsoiusagfoihsaofhai"
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

	for i, val := range post {
		fmt.Println(i, " => ")
		fmt.Println("\t title : ", val.Title)
		fmt.Println("\t category : ", val.Category.Name)
		fmt.Println("\t created_at :", val.CreatedAt)
	}
}

func ShowDetailePostControllers(id int) {
	var post models.Post
	err := DB.Preload("Category").Preload("User").Find(&post, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	fmt.Println("\t UserID : ", post.User.FirstName)
	fmt.Println("\t Title : ", post.Title)
	fmt.Println("\t CategoryID : ", post.Category.Name)
	fmt.Println("\t Img : ", post.Img)
	fmt.Println("\t Description : ", post.Description)
}

func UpdatePostControllers(id int) {
	var post models.Post

	err := DB.Find(&post, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	post.Title = "change"
	post.Img = "change"
	post.Description = "change"

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
