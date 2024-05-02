package controllers

import (
	"log"
	"strconv"

	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/gin-gonic/gin"
)
//Done 
func GetAllMenu(c *gin.Context) {
	var menu []models.Menu
	if err := initializers.DB.Unscoped().Table("menu_items").Find(&menu).Error; err != nil {
		log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": menu})
}
//Done
func CreateMenuItem(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBind(&menu); err != nil {
		log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := initializers.DB.Table("menu_items").Create(&menu).Error; err != nil {
		log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success create"})
}
//Done
func UpdateMenuItem(c *gin.Context) {
	id := c.Param("id")

	var menuRequest models.Menu
	if err := c.ShouldBind(&menuRequest); err != nil {
		log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	menuRequest.ID , _ = strconv.Atoi(id)

    if err := initializers.DB.Table("menu_items").Save(&menuRequest).Error; err != nil {
		log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
		return
    }
	c.JSON(200, gin.H{"message": "Success update"})
}
//MakeToFIX
func DeleteMenuItem(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	if err := initializers.DB.First(&menu, id).Error; err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := initializers.DB.Table("menu_items").Delete(&menu).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success delete"})
}

