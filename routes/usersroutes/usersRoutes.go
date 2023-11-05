package usersroutes

import (
	"github.com/dhanarrizky/go-blog/middleware"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("users")
	r.GET("users/:id")
	r.PUT("users/:id")    // for updatred the account
	r.DELETE("users/:id") // for delete the account
	CategoryRoutes(r)
	PostRoutes(r)
}

func CategoryRoutes(r *gin.Engine) {
	r.GET("users/category")
	r.POST("users/category")
	r.PUT("users/category")
	r.DELETE("users/category")
}

func PostRoutes(r *gin.Engine) {
	r.GET("users/Post")
	r.POST("users/Post")
	r.PUT("users/Post/:id")
	r.DELETE("users/Post/:id")
}
