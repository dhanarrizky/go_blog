package routes

import (
	"github.com/dhanarrizky/go-blog/controllers"
	"github.com/dhanarrizky/go-blog/middleware"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("users", controllers.ShowAllUserControllers())
	r.GET("users/:id", controllers.ShowUserDetaileControllers())
	r.PUT("users/:id", controllers.UpdateUserControllers())    // for updatred the account
	r.DELETE("users/:id", controllers.DeleteUserControllers()) // for delete the account
	CategoryRoutes(r)
	PostRoutes(r)
}

func CategoryRoutes(r *gin.Engine) {
	r.GET("users/category")
	r.POST("users/category")
	r.PUT("users/category/:id")
	r.DELETE("users/category/:id")
}

func PostRoutes(r *gin.Engine) {
	r.GET("users/Post")
	r.POST("users/Post")
	r.PUT("users/Post/:id")
	r.DELETE("users/Post/:id")
}
