// go_blog

package main

import (
	"fmt"

	"github.com/dhanarrizky/go-blog/controllers"
	"github.com/dhanarrizky/go-blog/database"
)

func main() {
	database.GetDB()
	mainMenu()
}

func mainMenu() {
	var inp int
	fmt.Println("============================== Main Menu ==============================")
	fmt.Println("choice \n" +
		"1. users\n" +
		"2. category\n" +
		"3. post\n")
	fmt.Scan(&inp)
	if inp == 1 {
		users()
	} else if inp == 2 {
		category()
	} else if inp == 3 {
		post()
	}
}

func users() {
	fmt.Println("============================== users ==============================")
	var inp, inp2 int

	fmt.Println("choice \n" +
		"1. create\n" +
		"2. showAll\n" +
		"3. showOne\n" +
		"4. update\n" +
		"5. delete\n" +
		"6. main menu")
	fmt.Scan(&inp)
	if inp == 1 {
		controllers.CreateUsersControllers()
	} else if inp == 2 {
		controllers.ShowAllUserControllers()
	} else if inp == 3 {
		controllers.ShowAllUserControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.ShowDetaileControllers(inp2)
	} else if inp == 4 {
		controllers.ShowAllUserControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.UpdateUserControllers(inp2)
	} else if inp == 5 {
		controllers.ShowAllUserControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.DeleteUserControllers(inp2)
	} else {
		mainMenu()
	}
}

func category() {
	fmt.Println("============================== category ==============================")
	var inp, inp2 int

	fmt.Println("choice \n" +
		"1. create\n" +
		"2. showAll\n" +
		"3. showOne\n" +
		"4. update\n" +
		"5. delete\n" +
		"6. main menu")
	fmt.Scan(&inp)
	if inp == 1 {
		controllers.CreateCategoryControllers()
	} else if inp == 2 {
		controllers.ShowAllCategoryControllers()
	} else if inp == 3 {
		controllers.ShowAllCategoryControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.ShowDetaileCategoryControllers(inp2)
	} else if inp == 4 {
		controllers.ShowAllCategoryControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.UpdateCategoryControllers(inp2)
	} else if inp == 5 {
		controllers.ShowAllCategoryControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.DeleteCategoryControllers(inp2)
	} else {
		mainMenu()
	}
}

func post() {
	fmt.Println("============================== post ==============================")
	var inp, inp2 int

	fmt.Println("choice \n" +
		"1. create\n" +
		"2. showAll\n" +
		"3. showOne\n" +
		"4. update\n" +
		"5. delete\n" +
		"6. main menu")
	fmt.Scan(&inp)
	if inp == 1 {
		controllers.CreatePostControllers()
	} else if inp == 2 {
		controllers.ShowAllPostControllers()
	} else if inp == 3 {
		controllers.ShowAllPostControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.ShowDetailePostControllers(inp2)
	} else if inp == 4 {
		controllers.ShowAllPostControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.UpdatePostControllers(inp2)
	} else if inp == 5 {
		controllers.ShowAllPostControllers()
		fmt.Print("choice One Id: ")
		fmt.Scan(&inp2)
		controllers.DeletePostControllers(inp2)
	} else {
		mainMenu()
	}
}
