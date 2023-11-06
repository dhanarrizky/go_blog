package routes

import (
	"github.com/dhanarrizky/go-blog/controllers"
	"github.com/gin-gonic/gin"
)

func UsersAuthencticationRoutes(r *gin.Engine) {
	r.POST("users/signup", controllers.UsersSignup())
	r.POST("users/login", controllers.UsersLogin())
}
