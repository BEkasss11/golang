package routes

import (
	"github.com/BEkasss11/golang/controllers"
	"github.com/BEkasss11/golang/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signUp", middleware.SignUp)
	routes.POST("users/signIn", middleware.SignIn)
	routes.GET("users/validate", middleware.RequireAuth, middleware.ValidateUser)
	routes.POST("users/signOut", middleware.SignOut)
}

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users", controllers.ListOfUsers)
	routes.GET("/users/:id", controllers.GetUserByID)
	routes.DELETE("/users/:id", controllers.DeleteUserByID)
}

func MenuRoutes(routes *gin.Engine) {
	routes.GET("/menu", controllers.GetAllMenu)
	
}
