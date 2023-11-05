package usersroutes

import "github.com/gin-gonic/gin"

func UsersAuthencticationRoutes(r *gin.Engine) {
	r.POST("users/signup")
	r.POST("users/login")
}
