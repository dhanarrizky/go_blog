// go_blog

package main

import (
	"log"
	"os"

	"github.com/dhanarrizky/go-blog/database"
	"github.com/dhanarrizky/go-blog/routes"
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
	routes.UsersAuthencticationRoutes(route)
	routes.UsersRoutes(route)

	route.Run(":" + port)
}
