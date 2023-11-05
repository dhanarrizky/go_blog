// go_blog

package main

import (
	"log"
	"os"

	"github.com/dhanarrizky/go-blog/database"
	"github.com/dhanarrizky/go-blog/routes"
	usersroutes "github.com/dhanarrizky/go-blog/routes/usersRoutes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Run The Application...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.GetDB()

	route := gin.Default()
	route.Use(gin.Logger())
	routes.GuestRoutes(route)
	usersroutes.UsersAuthencticationRoutes(route)
	usersroutes.UsersRoutes(route)

	route.Run(":" + port)
}
