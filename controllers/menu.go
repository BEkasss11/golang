package controllers

import (
	"fmt"

	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllMenu(c *gin.Context) {
	var menu []models.Menu
	initializers.DB.Find(&menu)
	fmt.Println(menu)
}

func CreateMenuItem(c *gin.Context) {
	var menu []models.Menu
	initializers.DB.Find(&menu)
	fmt.Println(menu)
}