package controllers

import (
	"fmt"

	"github.com/dhanarrizky/go-blog/models"
)

func CreateCategoryControllers() {
	var Categories models.Categories

	err := DB.Create(&Categories)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("created category has been successfully")
	}

}

func ShowAllCategoryControllers() {
	var Categories []models.Categories

	err := DB.Find(&Categories)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	// for i, val := range Categories {
	// 	fmt.Println(i, " => ")
	// 	fmt.Println("\t Name of category: ", val.Name)
	// }
}

func ShowDetaileCategoryControllers(id int) {
	var Categories models.Categories

	err := DB.Find(&Categories, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	// fmt.Println("category name : ", Categories.Name)
}

func UpdateCategoryControllers(id int) {
	var Categories models.Categories

	err := DB.Find(&Categories, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	err = DB.Save(&Categories)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("updated category has been successfully")
	}
}

func DeleteCategoryControllers(id int) {
	var Categories models.Categories

	err := DB.Find(&Categories, id)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	err = DB.Delete(&Categories)
	if err.Error != nil {
		panic(err.Error.Error())
	}

	if err.RowsAffected > 0 {
		fmt.Println("deleted category has been successfully")
	}
}
