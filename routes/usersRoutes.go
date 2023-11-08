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
	r.POST("users/category", controllers.CreateCategoryControllers())
	r.GET("users/category", controllers.ShowAllCategoryControllers())
	r.GET("users/category/:id", controllers.ShowDetaileCategoryControllers())
	r.PUT("users/category/:id", controllers.UpdateCategoryControllers())
	r.DELETE("users/category/:id", controllers.DeleteCategoryControllers())
}

func PostRoutes(r *gin.Engine) {
	r.POST("users/post", controllers.CreatePostControllers())
	r.GET("users/post", controllers.ShowAllPostControllers())
	r.GET("users/post/:id", controllers.ShowAllPostControllers())
	r.PUT("users/post/:id", controllers.UpdatePostControllers())
	r.DELETE("users/post/:id", controllers.DeletePostControllers())
}
