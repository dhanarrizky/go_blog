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

	database.GetDB()
	route := gin.Default()
	// route.ForwardedByClientIP = true
	// route.SetTrustedProxies([]string{"127.0.0.1:9000", "192.168.1.2", "10.0.0.0/8"})

	if port == "" {
		port = "8080"
	}

	route.Use(gin.Logger())
	routes.GuestRoutes(route)
	routes.UsersAuthencticationRoutes(route)
	routes.UsersRoutes(route)

	route.Run(":" + port)
}
