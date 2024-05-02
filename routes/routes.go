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
	routes.POST("/menu/item", controllers.CreateMenuItem)
	routes.PUT("/menu/item/:id", controllers.UpdateMenuItem)
	routes.DELETE("/menu/item/:id", controllers.DeleteMenuItem)
}

func OrderRoutes(routes *gin.Engine) {
    routes.POST("/orders", controllers.CreateOrder)
    routes.GET("/orders/:id", controllers.GetOrderByID)
    // routes.PUT("/orders/:id", controllers.UpdateOrder)
    routes.DELETE("/orders/:id", controllers.DeleteOrder)

    routes.POST("/order_items", controllers.CreateOrderItem)
    routes.GET("/order_items/:id", controllers.GetOrderItemByID)
    // routes.PUT("/order_items/:id", controllers.UpdateOrderItem)
    routes.DELETE("/order_items/:id", controllers.DeleteOrderItem)
}


// func PaymentTypeRoutes(router *gin.Engine) {
//     router.GET("/payment_types", controllers.GetAllPaymentTypes)
//     router.POST("/payment_types", controllers.CreatePaymentType)
//     router.GET("/payment_types/:id", controllers.GetPaymentTypeByID)
//     router.PUT("/payment_types/:id", controllers.UpdatePaymentTypeByID)
//     router.DELETE("/payment_types/:id", controllers.DeletePaymentTypeByID)
// }