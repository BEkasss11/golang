package controllers

import (
	"log"
	"strconv"
	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/gin-gonic/gin"
)

// Done
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBind(&order); err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := initializers.DB.Table("orders").Create(&order).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success create"})
}

// Done
func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := initializers.DB.Unscoped().Table("orders").First(&order, c.Param("id")).Error; err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": order})
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var orderRequest models.Order
	if err := c.ShouldBind(&orderRequest); err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	parsedId, _ := strconv.Atoi(id)

	orderRequest.ID = uint(parsedId)

	if err := initializers.DB.Table("orders").Save(&orderRequest).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success update"})
}

// Maketofix
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := initializers.DB.First(&order, id).Error; err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&order).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success delete"})
}

// Done
func CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBind(&orderItem); err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := initializers.DB.Table("order_items").Create(&orderItem).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success create"})
}

// Done
func GetOrderItemByID(c *gin.Context) {
	var orderItem models.OrderItem
	if err := initializers.DB.Unscoped().Table("order_items").First(&orderItem, c.Param("id")).Error; err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": orderItem})
}

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")

	var orderItemRequest models.OrderItem
	if err := c.ShouldBind(&orderItemRequest); err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	parsedId, _ := strconv.Atoi(id)

	orderItemRequest.ID = uint(parsedId)

	if err := initializers.DB.Save(&orderItemRequest).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success update"})
}

// Done
func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem
	if err := initializers.DB.First(&orderItem, id).Error; err != nil {
		log.Println(1)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := initializers.DB.Table("order_items").Delete(&orderItem).Error; err != nil {
		log.Println(2)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success delete"})
}
