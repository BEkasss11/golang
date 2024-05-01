package main

import (	
	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	

	routes.MenuRoutes(router)

	router.Run(":1232")
}
