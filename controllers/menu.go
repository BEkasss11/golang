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
	var menu models.Menu
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(400, gin.H{"message": "ERROR_CREATE_ITEM"})
		return
	}
	if err := initializers.DB.Create(&menu).Error; err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success"})
}

func UpdateMenuItem(c *gin.Context) {
	var menuRequest models.Menu
	if err := c.ShouldBind(&menuRequest); err != nil {
		c.JSON(400, gin.H{"message": "ERROR_CREATE_ITEM"})
		return
	}
	var menu models.Menu
	if err := initializers.DB.First(&menu, id).Error; err != nil {
        return err
    }

	if err := initializers.DB.Create(&menu).Error; err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success"})
}

func DeleteMenuItem(c *gin.Context) {
}
