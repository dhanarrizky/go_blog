package routes

import "github.com/gin-gonic/gin"

func GuestRoutes(r *gin.Engine) {
	r.GET("/")
	r.GET("/users-profile/:id")
	r.GET("/users-profile/post")
	r.GET("/users-profile/post/detaile:id")
	r.GET("/category")
}
