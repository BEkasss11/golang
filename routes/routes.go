package routes

import (
	"github.com/BEkasss11/golang/controllers"
	"github.com/BEkasss11/golang/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("/users/signUp", middleware.SignUp)
	routes.POST("users/signIn", middleware.SignIn)
	routes.GET("users/validate", middleware.RequireAuth, middleware.ValidateUser)
	routes.POST("users/signOut", middleware.SignOut)
}

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users",middleware.RequireAuth, middleware.IsAdmin, controllers.ListOfUsers)
	routes.GET("/users/:id",middleware.RequireAuth,middleware.IsAdmin, controllers.GetUserByID)
	routes.DELETE("/users/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.DeleteUserByID)
}

func MenuRoutes(routes *gin.Engine) {
	routes.GET("/menu", middleware.RequireAuth,controllers.GetAllMenu)
	routes.POST("/menu/item",middleware.RequireAuth, controllers.CreateMenuItem)
	routes.PUT("/menu/item/:id", middleware.RequireAuth, middleware.IsAdmin, controllers.UpdateMenuItem)
	routes.DELETE("/menu/item/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.DeleteMenuItem)
}

func OrderRoutes(routes *gin.Engine) {
    routes.POST("/orders",middleware.RequireAuth, controllers.CreateOrder)
    routes.GET("/orders/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.GetOrderByID)
    routes.PUT("/orders/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.UpdateOrder)
    routes.DELETE("/orders/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.DeleteOrder)

    routes.POST("/order_items",middleware.RequireAuth, controllers.CreateOrderItem)
    routes.GET("/order_items/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.GetOrderItemByID)
    routes.PUT("/order_items/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.UpdateOrderItem)
    routes.DELETE("/order_items/:id",middleware.RequireAuth, middleware.IsAdmin, controllers.DeleteOrderItem)
}
